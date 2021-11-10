package main

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoLog struct {
	Request_number int    `json:"request_number"`
	Gameid         int    `json:"gameid"`
	Gamename       string `json:"gamename"`
	Winner         string `json:"winner"`
	Players        int    `json:"players"`
	Worker         string `json:"worker"`
}

const (
	MONGODB_URI = "mongodb://admin:H3XT3tpQ3KeLTPQ8@35.188.126.89:27017"
	MONGO_DB    = "squidgame"
	MONGO_COL   = "logs"
)

func connectMongo(ctx context.Context) (*mongo.Collection, error) {
	/* Connect to my cluster */
	mongoClient, err := mongo.NewClient(options.Client().ApplyURI(MONGODB_URI))
	if err != nil {
		return nil, err
	}
	err = mongoClient.Connect(ctx)
	if err != nil {
		return nil, err
	}
	// defer mongoClient.Disconnect(ctx)
	if err := mongoClient.Ping(ctx, readpref.Primary()); err != nil {
		return nil, err
	}
	db := mongoClient.Database(MONGO_DB)
	col := db.Collection(MONGO_COL)
	return col, nil
}

func insertMongoLog(l MongoLog, col *mongo.Collection, ctx context.Context) error {
	_, insertErr := col.InsertOne(ctx, l)
	if insertErr != nil {
		return insertErr
	}
	// fmt.Println("Mongo ok")

	return nil
}
