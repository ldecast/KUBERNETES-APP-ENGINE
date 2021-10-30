package main

import (
	"context"
	game "go/games"
	"log"

	"google.golang.org/grpc"
)

func main() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := game.NewGameServiceClient(conn)

	g := game.Request{Gameid: 1, Gamename: "ldecast", Players: 55, RequestNumber: 2589}

	response, err := c.Play(context.Background(), &game.ServerRequest{Request: &g})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s", response)

}
