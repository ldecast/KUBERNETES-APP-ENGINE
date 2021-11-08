package main

import (
	kafkaconsumer "worker/kafka"

	"github.com/go-redis/redis"
)

const REDIS_HOST = "35.193.105.209"
const REDIS_PORT = "6379"
const REDIS_PASS = "Y5hNsA9hCnvDXXQLUjFuQxU3KKtwHrXW"

func connectRedis() (*redis.Client, error) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     REDIS_HOST + ":" + REDIS_PORT,
		Password: REDIS_PASS,
		DB:       0,
	})
	_, err := redisClient.Ping().Result()
	if err != nil {
		return nil, err
	}
	return redisClient, nil
}

func updateRedisValues(l kafkaconsumer.Log, client *redis.Client) error {
	// setting with `Key` and `Value` with no expiration time (0)`.
	client.Set("last_request", l.Request_number, 0)
	client.Set("last_gameid", l.Gameid, 0)
	client.Set("last_gamename", l.Gamename, 0)
	client.Set("last_winner", l.Winner, 0)
	client.Set("last_players", l.Players, 0)
	client.Set("last_worker", l.Worker, 0)
	// fmt.Println("Redis ok")

	return nil
}
