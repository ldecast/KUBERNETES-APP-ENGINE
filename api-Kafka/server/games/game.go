package games

import (
	context "context"

	"fmt"
	"log"
	"math/rand"
	"strconv"
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

func publish(msg Log) error {
	/* Publicar en Kafka */
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
		Worker:         "Kafka"}
	/* Insertar a cola de Kafka */
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
		Worker:         "Kafka"}
	/* Insertar a cola de Kafka */
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
		Worker:         "Kafka"}
	/* Insertar a cola de Kafka */
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
