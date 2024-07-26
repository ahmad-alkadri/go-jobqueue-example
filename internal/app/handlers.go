package app

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
)

func HandleProcess(jobQueue *JobQueue) func (w http.ResponseWriter, r *http.Request)  {
	return func(w http.ResponseWriter, r *http.Request) {
		jobID := uuid.New().String()
		job := Job{ID: jobID}
	
		if err := jobQueue.AddJob(job); err != nil {
			http.Error(w, "Failed to enqueue job", http.StatusInternalServerError)
			return
		}
	
		w.WriteHeader(http.StatusAccepted)
		response := map[string]string{"jobID": jobID}
		json.NewEncoder(w).Encode(response)
	}
}

func HandleStatus(jobQueue *JobQueue) func (w http.ResponseWriter, r *http.Request)  {
	return func(w http.ResponseWriter, r *http.Request) {
		jobID := r.URL.Query().Get("id")
		job, err := jobQueue.GetJob(jobID)
		if err != nil {
			http.NotFound(w, r)
			return
		}
		json.NewEncoder(w).Encode(job)
	}
}
