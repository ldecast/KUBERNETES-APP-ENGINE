package games

import (
	context "context"
	"fmt"
	"log"
	"math/rand"
	kafkaproducer "server/kafka"
	"strconv"
	"sync/atomic"
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

var brokers = [...]string{""}

const topic = "first_kafka_topic"

func publish(msg Log, publisher kafkaproducer.Publisher) error {
	if err := publisher.Publish(context.Background(), msg); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func MaxPlayer(players int) (string, string) {
	return "MaxPlayer", strconv.Itoa(players)
}

func MinPlayer(players int) (string, string) {
	return "MinPlayer", "1"
}

func RandomPlayer(players int) (string, string) {
	randomIndex := rand.Intn(players)
	if randomIndex == 0 {
		if players > 1 {
			randomIndex = 2
		} else {
			randomIndex = 1
		}
	}
	return "RandomPlayer", strconv.Itoa(randomIndex)
}

var request_number int64 = 0

var publisher = kafkaproducer.NewPublisher(brokers[:], topic)

func (s *Server) Play(ctx context.Context, in *ServerRequest) (*ServerResponse, error) {
	game := in.Request
	log.Printf("Receive request from client: %s", game)

	/* Crear el log */
	l := Log{
		Request_number: int(atomic.AddInt64(&request_number, 1)),
		Gameid:         int(game.Gameid),
		Gamename:       "",
		Winner:         "",
		Players:        int(game.Players),
		Worker:         "Kafka"}

	switch l.Gameid {
	case 1:
		l.Gamename, l.Winner = MaxPlayer(l.Players)
	case 2:
		l.Gamename, l.Winner = MinPlayer(l.Players)
	case 3:
		l.Gamename, l.Winner = RandomPlayer(l.Players)
	default:
		fmt.Println("No existe ningún juego con ese identificador, intente de nuevo.")
		return &ServerResponse{Status: "[ERR - 400]"}, nil
	}

	/* Publicar en Kafka */
	err := publish(l, publisher)
	if err != nil {
		return &ServerResponse{Status: "[ERR - 400]"}, nil
	}

	return &ServerResponse{Status: "[OK - 200]"}, nil
}
