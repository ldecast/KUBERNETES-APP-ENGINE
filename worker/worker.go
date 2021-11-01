package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"cloud.google.com/go/pubsub"
	"github.com/go-redis/redis"
	"go.mongodb.org/mongo-driver/mongo"
)

const projectID = "pubsub-tester-330701"
const subID = "logs-sub"

func subscribe(col *mongo.Collection, redisClient *redis.Client, ctx context.Context) error {
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "key.json")

	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		fmt.Println(err)
		// return fmt.Errorf("pubsub.NewClient: %v", err)
	}
	defer client.Close()

	sub := client.Subscription(subID)

	var l MongoLog
	err = sub.Receive(ctx, func(ctx context.Context, m *pubsub.Message) {
		log.Printf("Got message: %s", m.Data)
		json.Unmarshal(m.Data, &l)
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
		/* Confirmar mensaje de PubSub */
		m.Ack()
	})
	if err != nil {
		fmt.Println(err)
		// return err
	}

	return nil
}

func main() {
	fmt.Println("Go Pub/Sub worker!")
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
	/* Suscribirse a la cola de mensajes de Google PubSub */
	err = subscribe(mongoClient, redisClient, ctx)
	if err != nil {
		log.Fatal(err)
	}
}
