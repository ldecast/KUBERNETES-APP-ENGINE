package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/go-redis/redis"
	"github.com/streadway/amqp"
	"go.mongodb.org/mongo-driver/mongo"
)

const queue_name = "RABBITMQ_QUEUE"

// "amqp://guest:guest@localhost:5672/rabbit"
var rabbitmq_host string
var rabbitmq_port string

func subscribe(col *mongo.Collection, redisClient *redis.Client, ctx context.Context) error {
	/* Suscribirse a RabbitMQ */
	rabbit_connection, err := amqp.Dial("amqp://guest:guest@" + rabbitmq_host + ":" + rabbitmq_port + "/")
	if err != nil {
		fmt.Println("Failed Initializing RabbitMQ Broker Connection")
		return err
	}

	ch, err := rabbit_connection.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	if err != nil {
		return err
	}

	msgs, err := ch.Consume(
		queue_name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	forever := make(chan bool)
	var l MongoLog
	go func() {
		for d := range msgs {
			fmt.Printf("Recieved RabbitMQ Message: %s\n", d.Body)

			json.Unmarshal(d.Body, &l)
			/* Insertar a Mongo */
			err_mongo := insertMongoLog(l, col, ctx)
			if err_mongo != nil {
				fmt.Println(err_mongo)
			}
			/* Actualizar metadata en Redis */
			err_redis := updateRedisValues(l, redisClient)
			if err_redis != nil {
				fmt.Println(err_redis)
			}
		}
	}()

	fmt.Println("Worker connected to RabbitMQ Instance")
	fmt.Println(" [*] - Waiting for messages")
	<-forever
	return nil
}

func main() {
	fmt.Println("Go RabbitMQ worker!")

	rabbitmq_host = os.Getenv("rabbitmq_host")
	if rabbitmq_host == "" {
		rabbitmq_host = "localhost"
		// log.Fatal("rabbitmq_host is not defined as environment variable")
	}

	rabbitmq_port = os.Getenv("rabbitmq_port")
	if rabbitmq_port == "" {
		rabbitmq_port = "5672"
		// log.Fatal("rabbitmq_port is not defined as environment variable")
	}

	ctx := context.Background()
	/* Conectar a Mongo y obtener la colección donde se inserta cada log */
	mongoClient, err := connectMongo(ctx)
	if err != nil {
		log.Fatal(err)
	}
	/* Conectar a Redis y obtener el cliente donde se actualiza cada valor */
	redisClient, err := connectRedis()
	if err != nil {
		log.Fatal(err)
	}
	/* Suscribirse a la cola de mensajes de RabbitMQ */
	err = subscribe(mongoClient, redisClient, ctx)
	if err != nil {
		log.Fatal(err)
	}
}
