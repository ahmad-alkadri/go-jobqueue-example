package app

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx context.Context = context.Background()

// Worker represents a worker that processes jobs.
type Worker struct {
	ID         int
	jobQueue   *JobQueue
	quit       chan bool
}

// NewWorker creates a new Worker.
func NewWorker(id int, jobQueue *JobQueue) *Worker {
	return &Worker{
		ID:         id,
		jobQueue:   jobQueue,
		quit:       make(chan bool),
	}
}

// Start starts the worker to process jobs.
func (w *Worker) Start() {
	go func() {
		for {
			select {
			case <-w.quit:
				return
			default:
				jobID, err := w.jobQueue.redisClient.BRPop(ctx, 1*time.Second, "jobQueue").Result()
				if err != nil {
					if err != redis.Nil {
						fmt.Printf("Worker %d encountered error: %v\n", w.ID, err)
					}
					continue
				}
				if len(jobID) < 2 {
					continue
				}
				id := jobID[1]
				job, err := w.jobQueue.GetJob(id)
				if err != nil {
					fmt.Printf("Worker %d encountered error: %v\n", w.ID, err)
					continue
				}
				fmt.Printf("Worker %d started job %s\n", w.ID, job.ID)
				job.Status = "Processing"
				w.jobQueue.UpdateJob(job)

				// Simulate long processing time
				time.Sleep(10 * time.Second)

				job.Status = "Completed"
				job.Result = "Processed data for job " + job.ID
				w.jobQueue.UpdateJob(job)
				fmt.Printf("Worker %d completed job %s\n", w.ID, job.ID)
			}
		}
	}()
}


// Stop stops the worker.
func (w *Worker) Stop() {
	go func() {
		w.quit <- true
	}()
}

// StartWorkerPool starts a pool of workers.
func StartWorkerPool(queue *JobQueue, numWorkers int) []*Worker {
	var workers []*Worker
	for i := 1; i <= numWorkers; i++ {
		worker := NewWorker(i, queue)
		worker.Start()
		workers = append(workers, worker)
	}
	return workers
}