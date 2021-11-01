package main

import (
	"fmt"

	"github.com/go-redis/redis"
)

const REDIS_HOST = "127.0.0.1"
const REDIS_PORT = "6379"

func updateRedisValues(l MongoLog) error {
	client := redis.NewClient(&redis.Options{
		Addr:     REDIS_HOST + ":" + REDIS_PORT,
		Password: "Y5hNsA9hCnvDXXQLUjFuQxU3KKtwHrXW",
		DB:       0,
	})
	_, err := client.Ping().Result()
	if err != nil {
		return err
	}
	// setting with `Key` and `Value` with no expiration time (0)`.
	client.Set("last_request", l.Request_number, 0)
	client.Set("last_gameid", l.Gameid, 0)
	client.Set("last_gamename", l.Gamename, 0)
	client.Set("last_winner", l.Winner, 0)
	client.Set("last_players", l.Players, 0)
	client.Set("last_worker", l.Worker, 0)
	fmt.Println("Redis ok")
	return nil
}
