package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"server/games"

	"github.com/streadway/amqp"
	"google.golang.org/grpc"
)

const server_PORT = 9000

var address_rabbitmq string // = "amqp://guest:guest@localhost:5672/rabbit"

func initRabbit() error {
	var errConn error
	games.Rabbit_connection, errConn = amqp.Dial("amqp://guest:guest@" + address_rabbitmq + ":5672/")
	if errConn != nil {
		fmt.Println("Failed Initializing RabbitMQ Broker Connection")
		return errConn
	} else {
		fmt.Println("Publisher connected to RabbitMQ Instance")
	}

	// Let's start by opening a channel to our RabbitMQ instance
	ch, err := games.Rabbit_connection.Channel()
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer ch.Close()

	// with the instance and declare Queues that we can publish and
	_, err = ch.QueueDeclare(
		games.QueueName,
		true,
		false,
		false,
		false,
		nil,
	)
	// fmt.Println(q)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func main() {

	fmt.Println("Go gRPC Server for API RabbitMQ!")

	address_rabbitmq = os.Getenv("address_rabbitmq")
	if address_rabbitmq == "" {
		log.Fatal("address_rabbitmq is not defined as environment variable")
	}

	// Iniciar conexión con RabbitMQ
	err := initRabbit()
	if err != nil {
		log.Fatal(err)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", server_PORT))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	} /* else {
		fmt.Println("gRPC Server ok")
	} */

	s := games.Server{}

	grpcServer := grpc.NewServer()

	games.RegisterGameServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
