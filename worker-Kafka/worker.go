package main

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/mongo"
)

func subscribe(col *mongo.Collection, redisClient *redis.Client, ctx context.Context) error {
	/* Suscribirse a Kafka */

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
