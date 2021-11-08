package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	kafkaconsumer "worker/kafka"

	"github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	// host    = os.Getenv("HOST")
	brokers = ""
	topic   = "first_kafka_topic"
)

func subscribe(col *mongo.Collection, redisClient *redis.Client, ctx context.Context) error {
	/* Suscribirse a Kafka */
	chMsg := make(chan kafkaconsumer.Log)
	chErr := make(chan error)
	consumer := kafkaconsumer.NewConsumer(strings.Split(brokers, ","), topic)

	go func() {
		consumer.Read(context.Background(), chMsg, chErr)
	}()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	for {
		select {
		case <-quit:
			goto end
		case l := <-chMsg:
			log.Printf("Got message: %+v\n", l)
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
		case err := <-chErr:
			log.Println(err)
		}
	}
end:
	fmt.Println("\nyou have abandoned the Kafka topic")
	return nil
}

func main() {
	fmt.Println("Go Kafka worker!")
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
	/* Suscribirse a la cola de mensajes de Kafka */
	err = subscribe(mongoClient, redisClient, ctx)
	if err != nil {
		log.Fatal(err)
	}
}
