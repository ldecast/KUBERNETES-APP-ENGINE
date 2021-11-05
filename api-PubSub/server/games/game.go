package games

import (
	context "context"
	"encoding/json"
	"os"

	"fmt"
	"log"
	"math/rand"
	"strconv"

	"cloud.google.com/go/pubsub"
)

type Server struct {
}

type Log struct {
	Request_number int    `json:"request_number"`
	Gameid         int    `json:"gameid"`
	Gamename       string `json:"gamename"`
	Winner         string `json:"winner"`
	Players        int    `json:"players"`
	Worker         string `json:"worker"`
}

const projectID = "inner-bridge-331015"
const topicID = "squidgame-topic"
const subID = "squidgame-subs"

func publish(msg Log) error {
	/* Publicar en Pub/Sub */
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "key.json")
	os.Setenv("PROJECT", subID)

	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("pubsub.NewClient: %v", err)
	}
	defer client.Close()

	notif, err := json.Marshal(msg)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("json.Marshal: %v", err)
	}

	t := client.Topic(topicID)
	result := t.Publish(ctx, &pubsub.Message{
		Data: notif,
	})
	// ID is returned for the published message.
	id, err := result.Get(ctx)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Printf("Published a message; msg ID: %v\n", id)
	return nil
}

func MaxPlayer(game *Request) *ServerResponse {
	request_number++
	/* Crear el log */
	l := Log{
		Request_number: request_number,
		Gameid:         1,
		Gamename:       "MaxPlayer",
		Winner:         strconv.Itoa(int(game.Players)),
		Players:        int(game.Players),
		Worker:         "PubSub"}
	/* Insertar a cola de PubSub */
	err := publish(l)
	if err != nil {
		fmt.Printf("Error publishing log")
		return &ServerResponse{Status: "[ERR - 500]"}
	}
	return &ServerResponse{Status: "[OK - 200]"}
}

func MinPlayer(game *Request) *ServerResponse {
	request_number++
	/* Crear el log */
	l := Log{
		Request_number: request_number,
		Gameid:         2,
		Gamename:       "MinPlayer",
		Winner:         "1",
		Players:        int(game.Players),
		Worker:         "PubSub"}
	/* Insertar a cola de PubSub */
	err := publish(l)
	if err != nil {
		fmt.Printf("Error publishing log")
		return &ServerResponse{Status: "[ERR - 500]"}
	}
	return &ServerResponse{Status: "[OK - 200]"}
}

func RandomPlayer(game *Request) *ServerResponse {
	request_number++
	/* Crear el log */
	randomIndex := rand.Intn(int(game.Players))
	if randomIndex == 0 {
		if game.Players > 1 {
			randomIndex = 2
		} else {
			randomIndex = 1
		}
	}
	l := Log{
		Request_number: request_number,
		Gameid:         3,
		Gamename:       "RandomPlayer",
		Winner:         strconv.Itoa(randomIndex),
		Players:        int(game.Players),
		Worker:         "PubSub"}
	/* Insertar a cola de PubSub */
	err := publish(l)
	if err != nil {
		fmt.Printf("Error publishing log")
		return &ServerResponse{Status: "[ERR - 500]"}
	}
	return &ServerResponse{Status: "[OK - 200]"}
}

var request_number int = 0

func (s *Server) Play(ctx context.Context, in *ServerRequest) (*ServerResponse, error) {
	game := in.Request
	log.Printf("Receive request %d from client: %s", request_number+1, game)

	switch game.Gameid {
	case 1:
		return MaxPlayer(game), nil
	case 2:
		return MinPlayer(game), nil
	case 3:
		return RandomPlayer(game), nil
	default:
		fmt.Println("No existe ningún juego con ese identificador, intente de nuevo.")
		return &ServerResponse{Status: "[ERR - 400]"}, nil
	}
}
