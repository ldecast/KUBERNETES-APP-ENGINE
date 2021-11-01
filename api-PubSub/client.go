package main

import (
	"context"
	"encoding/json"
	"fmt"
	game "go/games"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/rs/cors"
	"google.golang.org/grpc"
)

type RequestBody struct {
	Gameid         int    `json:"gameid"`
	Gamename       string `json:"gamename"`
	Players        int    `json:"players"`
	Request_number int    `json:"request_number"`
}

func sayHi(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello from gRPC Pub/Sub producer!")
}

func sendRequest(w http.ResponseWriter, r *http.Request) {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()
	// defer r.Body.Close()

	c := game.NewGameServiceClient(conn)

	reqBody, _ := ioutil.ReadAll(r.Body)
	var request RequestBody
	json.Unmarshal(reqBody, &request)

	g := game.Request{Gameid: int32(request.Gameid), Gamename: request.Gamename, Players: int32(request.Players), RequestNumber: int32(request.Request_number)}

	response, err := c.Play(context.Background(), &game.ServerRequest{Request: &g})
	if err != nil {
		log.Fatalf("Error when calling Play: %s", err)
	}
	log.Printf("Response from gRPC server: %s", response.Status)
	fmt.Fprintf(w, "Response from gRPC server: %s", response.Status)
}

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", sayHi)
	mux.HandleFunc("/play", sendRequest)

	// cors.Default() setup the middleware with default options being
	// all origins accepted with simple methods (GET, POST).
	handler := cors.Default().Handler(mux)

	// Determine port for HTTP service.
	port := os.Getenv("PORT")
	if port == "" {
		port = "10000"
		log.Printf("defaulting to port %s", port)
	}

	// Start HTTP server.
	log.Printf("Go server listening on port: %s", port)
	if err := http.ListenAndServe(":"+port, handler); err != nil {
		log.Fatal(err)
	}
}
