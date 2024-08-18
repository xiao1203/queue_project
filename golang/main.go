package main

import (
	"context"
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
)

type Job struct {
	ID   string `json:"id"`
	Data string `json:"data"`
}

func main() {
	ctx := context.Background()
	redisAddr := os.Getenv("REDIS_URL")
	if redisAddr == "" {
		redisAddr = "redis:6379"
	}

	redisClient := redis.NewClient(&redis.Options{
		Addr: redisAddr,
	})

	for {
		job := Job{
			ID:   time.Now().String(),
			Data: "Some data to process",
		}

		jobJSON, err := json.Marshal(job)
		if err != nil {
			log.Printf("Error marshalling job: %v", err)
			continue
		}

		err = redisClient.RPush(ctx, "job_queue", jobJSON).Err()
		if err != nil {
			log.Printf("Error enqueueing job: %v", err)
		} else {
			log.Printf("Job enqueued: %s", job.ID)
		}

		time.Sleep(5 * time.Second)
	}
}
