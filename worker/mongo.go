package main

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const MONGODB_URI = "mongodb+srv://root:tZqqBxg6KnfQqhWA@cluster0.szsb6.mongodb.net/proyecto2-so1?retryWrites=true&w=majority"

type MongoLog struct {
	Request_number int    `json:"request_number"`
	Gameid         int    `json:"gameid"`
	Gamename       string `json:"gamename"`
	Winner         string `json:"winner"`
	Players        int    `json:"players"`
	Worker         string `json:"worker"`
}

func insertMongoLog(l MongoLog) error {
	/* Connect to my cluster */
	mongoClient, err := mongo.NewClient(options.Client().ApplyURI(MONGODB_URI))
	if err != nil {
		return err
	}
	ctx := context.Background()
	err = mongoClient.Connect(ctx)
	if err != nil {
		return err
	}
	defer mongoClient.Disconnect(ctx)
	if err := mongoClient.Ping(ctx, readpref.Primary()); err != nil {
		return err
	}
	db := mongoClient.Database("proyecto2-so1")
	col := db.Collection("squidgame")
	_, insertErr := col.InsertOne(ctx, l)
	if insertErr != nil {
		return insertErr
	} /* else {
		// get the inserted ID string
		newID := result.InsertedID
		fmt.Println("InsertOne() newID inserted:", newID)
		fmt.Println()
	} */
	fmt.Println("Mongo ok")

	return nil
}
