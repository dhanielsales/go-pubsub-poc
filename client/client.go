package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

var redisClient = redis.NewClient(&redis.Options{
	Addr: os.Getenv("REDIS_URL"),
})

const ACTION_KEY = "ACTION_KEY"

func main() {
	http.HandleFunc("/client", handler)
	fmt.Println("Client started at port 8080")
	fmt.Println("POST | http://localhost:8080/client")

	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		body, err := ioutil.ReadAll(r.Body)

		if err != nil {
			log.Printf("Error reading body: %v", err)
			http.Error(w, "can't read body", http.StatusBadRequest)
			return
		}

		if err := redisClient.Publish(ctx, ACTION_KEY, body).Err(); err != nil {
			panic(err)
		}

		fmt.Fprintf(w, "%s", body)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "{\"error\": {\"message\": \"Not found\"}}")
	}
}
