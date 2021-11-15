package main

import (
	"context"
	"fmt"
	"os"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

var redisClient = redis.NewClient(&redis.Options{
	Addr: os.Getenv("REDIS_URL"),
})

const ACTION_KEY = "ACTION_KEY"

func main() {
	fmt.Println("Worker started and waiting for messages")

	subscriber := redisClient.Subscribe(ctx, ACTION_KEY)

	for {
		msg, err := subscriber.ReceiveMessage(ctx)
		if err != nil {
			panic(err)
		}

		fmt.Println("Message Received: ", msg)
	}
}
