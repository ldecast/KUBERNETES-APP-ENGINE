package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"cloud.google.com/go/pubsub"
	"github.com/rs/cors"
)

func subscribe() error {
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "key.json")
	projectID := "pubsub-tester-330701"
	subID := "logs-sub"
	ctx := context.Background()
	client, err := pubsub.NewClient(ctx, projectID)
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("pubsub.NewClient: %v", err)
	}
	defer client.Close()

	sub := client.Subscription(subID)

	/* // Receive messages for 10 seconds.
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel() */

	/* // Create a channel to handle messages to as they come in.
	cm := make(chan *pubsub.Message)
	defer close(cm)

	// Handle individual messages in a goroutine.
	go func() {
		for msg := range cm {
			fmt.Printf("Got message :%q\n", string(msg.Data))
			fmt.Println("Attributes:")
			for key, value := range msg.Attributes {
				fmt.Printf("%s = %s", key, value)
			}
			msg.Ack()
		}
	}()

	// Receive blocks until the context is cancelled or an error occurs.
	err = sub.Receive(ctx, func(ctx context.Context, msg *pubsub.Message) {
		fmt.Println("ooo")
		cm <- msg
	})
	if err != nil {
		return fmt.Errorf("receive: %v", err)
	} */

	err = sub.Receive(ctx, func(ctx context.Context, m *pubsub.Message) {
		log.Printf("Got message: %s", m.Data)
		m.Ack()
	})
	if err != nil {
		fmt.Println(err)
		// Handle error.
	}

	return nil
}

/* func insert_Mongo() {

}

func insert_Redis() {

} */

func sayHi(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from Go Pub/Sub worker!")
}

func main() {
	fmt.Println("Go Pub/Sub worker!")
	err := subscribe()
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", sayHi)

	handler := cors.Default().Handler(mux)

	// Determine port for HTTP service.
	port := os.Getenv("PORT")
	if port == "" {
		port = "7878"
		log.Printf("defaulting to port %s", port)
	}

	// Start HTTP server.
	log.Printf("Go server listening on port: %s", port)
	if err := http.ListenAndServe(":"+port, handler); err != nil {
		log.Fatal(err)
	}
}
