package games

import (
	context "context"
	// "encoding/json"
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
	/* Publicar en RabbitMQ */
	return nil
}

func MaxPlayer(game *Request) *ServerResponse {
	/* Crear el log */
	l := Log{
		Request_number: int(game.RequestNumber),
		Gameid:         1,
		Gamename:       "MaxPlayer",
		Winner:         strconv.Itoa(int(game.Players)),
		Players:        int(game.Players),
		Worker:         "RabbitMQ"}
	/* Insertar a cola de RabbitMQ */
	err := publish(l)
	if err != nil {
		fmt.Printf("Error publishing log")
		return &ServerResponse{Status: "[ERR - 500]"}
	}
	return &ServerResponse{Status: "[OK - 200]"}
}

func MinPlayer(game *Request) *ServerResponse {
	/* Crear el log */
	l := Log{
		Request_number: int(game.RequestNumber),
		Gameid:         2,
		Gamename:       "MinPlayer",
		Winner:         "1",
		Players:        int(game.Players),
		Worker:         "RabbitMQ"}
	/* Insertar a cola de RabbitMQ */
	err := publish(l)
	if err != nil {
		fmt.Printf("Error publishing log")
		return &ServerResponse{Status: "[ERR - 500]"}
	}
	return &ServerResponse{Status: "[OK - 200]"}
}

func RandomPlayer(game *Request) *ServerResponse {
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
		Request_number: int(game.RequestNumber),
		Gameid:         3,
		Gamename:       "RandomPlayer",
		Winner:         strconv.Itoa(randomIndex),
		Players:        int(game.Players),
		Worker:         "RabbitMQ"}
	/* Insertar a cola de RabbitMQ */
	err := publish(l)
	if err != nil {
		fmt.Printf("Error publishing log")
		return &ServerResponse{Status: "[ERR - 500]"}
	}
	return &ServerResponse{Status: "[OK - 200]"}
}

func (s *Server) Play(ctx context.Context, in *ServerRequest) (*ServerResponse, error) {
	game := in.Request
	log.Printf("Receive message body from client: %s", game)

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
