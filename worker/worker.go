package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"cloud.google.com/go/pubsub"
)

const projectID = "pubsub-tester-330701"
const subID = "logs-sub"

func subscribe() error {
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "key.json")
	ctx := context.Background()
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
		err_mongo := insertMongoLog(l)
		if err_mongo != nil {
			fmt.Println(err_mongo)
		}
		/* Actualizar metadata en Redis */
		err_redis := updateRedisValues(l)
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
	err := subscribe()
	if err != nil {
		log.Fatal(err)
	}
}
