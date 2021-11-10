package main

import (
	"fmt"
	"log"
	"net"
	"server/games"
	"server/worker"

	"github.com/streadway/amqp"
	"google.golang.org/grpc"
)

const (
	server_PORT       = 9000
	string_Connection = "amqp://guest:guest@localhost:5672/rabbit"
)

func initRabbit() error {
	var errConn error
	games.Rabbit_connection, errConn = amqp.Dial(string_Connection)
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

	// Iniciar conexión con RabbitMQ
	err := initRabbit()
	if err != nil {
		log.Fatal(err)
	}

	// Iniciar escucha de mensajes
	go worker.StartWorker()

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
