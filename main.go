package main

import (
	"log"
	"net/http"

	app "github.com/ahmad-alkadri/go-jobqueue-example/internal/app"
	"github.com/go-redis/redis/v8"
)

func main() {
	var jobQueue *app.JobQueue
	// Connect to Redis
	redisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	defer redisClient.Close()

	jobQueue = app.NewJobQueue(redisClient)
	app.StartWorkerPool(jobQueue, 3)

	http.HandleFunc("/process", app.HandleProcess(jobQueue))
	http.HandleFunc("/status", app.HandleStatus(jobQueue))

	log.Println("Server is starting on port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}	
}
