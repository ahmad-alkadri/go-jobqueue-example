package app

import (
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v8"
)

// Job represents a unit of work.
type Job struct {
	ID     string `json:"id"`
	Status string `json:"status"`
	Result string `json:"result"`
}

// JobQueue represents a queue for jobs.
type JobQueue struct {
	redisClient *redis.Client
}

// NewJobQueue creates a new JobQueue.
func NewJobQueue(redisClient *redis.Client) *JobQueue {
	return &JobQueue{
		redisClient: redisClient,
	}
}

// AddJob adds a job to the queue with a TTL of 3 days.
func (jq *JobQueue) AddJob(job Job) error {
	job.Status = "Pending"
	jobData, err := json.Marshal(job)
	if err != nil {
		return err
	}
	if err := jq.redisClient.Set(ctx, job.ID, jobData, 72*time.Hour).Err(); err != nil {
		return err
	}
	return jq.redisClient.LPush(ctx, "jobQueue", job.ID).Err()
}

// GetJob retrieves a job by its ID.
func (jq *JobQueue) GetJob(id string) (Job, error) {
	jobData, err := jq.redisClient.Get(ctx, id).Result()
	if err != nil {
		return Job{}, err
	}
	var job Job
	if err := json.Unmarshal([]byte(jobData), &job); err != nil {
		return Job{}, err
	}
	return job, nil
}

// UpdateJob updates the status of a job and refreshes its TTL.
func (jq *JobQueue) UpdateJob(job Job) error {
	jobData, err := json.Marshal(job)
	if err != nil {
		return err
	}
	return jq.redisClient.Set(ctx, job.ID, jobData, 72*time.Hour).Err()
}