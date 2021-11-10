package games

import (
	context "context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"sync/atomic"

	"github.com/streadway/amqp"
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

var Rabbit_connection *amqp.Connection

const QueueName = "RABBITMQ_QUEUE"

func publish(msg Log) error {
	// Let's start by opening a channel to our RabbitMQ instance
	ch, err := Rabbit_connection.Channel()
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer ch.Close()

	body, err := json.Marshal(msg)
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = ch.Publish(
		"",
		QueueName,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		})
	if err != nil {
		log.Printf("Failed to publish a message: %s", err)
		return err
	} else {
		fmt.Println("Successfully Published Message to Queue")
		return nil
	}
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
		Worker:         "RabbitMQ"}

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

	/* Publicar en RabbitMQ */
	err := publish(l)
	if err != nil {
		return &ServerResponse{Status: "[ERR - 400]"}, nil
	}

	return &ServerResponse{Status: "[OK - 200]"}, nil
}
