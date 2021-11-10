package worker

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"server/games"

	"github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/mongo"
)

func subscribe(col *mongo.Collection, redisClient *redis.Client, ctx context.Context) error {
	/* Suscribirse a RabbitMQ */
	ch, err := games.Rabbit_connection.Channel()
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer ch.Close()

	if err != nil {
		fmt.Println(err)
		return err
	}

	msgs, err := ch.Consume(
		games.QueueName,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println(err)
		return err
	}

	forever := make(chan bool)
	var l games.Log
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

func StartWorker() {
	fmt.Println("Go RabbitMQ worker!")
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
