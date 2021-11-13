package main

import (
	"context"
	"encoding/json"
	"fmt"

	game "client/games"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/rs/cors"
	"google.golang.org/grpc"
)

type GameBody struct {
	Gameid   int    `json:"gameid"`
	Gamename string `json:"gamename"`
	Players  int    `json:"players"`
}

func sayHi(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from gRPC RabbitMQ producer!")
}

var gRPC_server_host string
var gRPC_server_port string

func sendRequest(w http.ResponseWriter, r *http.Request) {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial((gRPC_server_host + ":" + gRPC_server_port), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()
	// defer r.Body.Close()

	c := game.NewGameServiceClient(conn)

	reqBody, _ := ioutil.ReadAll(r.Body)
	var game_body GameBody
	json.Unmarshal(reqBody, &game_body)

	g := game.Request{
		Gameid:   int32(game_body.Gameid),
		Gamename: game_body.Gamename,
		Players:  int32(game_body.Players)}

	response, err := c.Play(context.Background(), &game.ServerRequest{Request: &g})
	if err != nil {
		log.Fatalf("Client error when calling Play: %s", err)
	}
	log.Printf("Response from gRPC server: %s", response.Status)
	fmt.Fprintf(w, "Response from gRPC server: %s", response.Status)
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", sendRequest)
	mux.HandleFunc("/hi", sayHi)

	// cors.Default() setup the middleware with default options being
	// all origins accepted with simple methods (GET, POST).
	handler := cors.Default().Handler(mux)

	gRPC_server_host = os.Getenv("server_host_rabbitmq")
	if gRPC_server_host == "" {
		log.Fatal("server_host_rabbitmq is not defined as environment variable")
	}

	gRPC_server_port = os.Getenv("server_port_rabbitmq")
	if gRPC_server_port == "" {
		log.Fatal("server_port_rabbitmq is not defined as environment variable")
	}

	// Determine port for HTTP service.
	port := os.Getenv("PORT")
	if port == "" {
		port = "10000"
		log.Printf("defaulting to port %s", port)
	}

	// Start HTTP server.
	log.Printf("Go client for RabbitMQ listening on port: %s", port)
	if err := http.ListenAndServe(":"+port, handler); err != nil {
		log.Fatal(err)
	}
}
