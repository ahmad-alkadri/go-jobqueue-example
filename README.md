
# go-jobqueue-example

This is an example implementation of a job queue using Go and Redis. Workers continuously check for new jobs in the queue, process them asynchronously, and store the results back in Redis. This setup is ideal for handling long-running tasks without blocking the main application.

## How to Clone

To clone the repository, use the following command:

```bash
git clone https://github.com/ahmad-alkadri/go-jobqueue-example.git
cd go-jobqueue-example
```

## How to Run

1. **Install Redis**: Ensure you have Redis installed and running on your local machine. If you have Docker, you can start Redis with:

   ```bash
   docker run --name redis -p 6379:6379 -d redis
   ```

2. **Install Dependencies**: Install the required Go dependencies:

   ```bash
   go mod tidy
   ```

3. **Run the Application**:

   ```bash
   go run main.go
   ```

## How to Test/Use This Example

1. **Submit a Job**: To submit a job, send a POST request to the `/process` endpoint. This will enqueue a new job.

   ```bash
   curl -X POST http://localhost:8080/process
   ```

   You will receive a response with a job ID:

   ```json
   {
     "jobID": "some-unique-job-id"
   }
   ```

2. **Check Job Status**: To check the status of a job, send a GET request to the `/status` endpoint with the job ID as a query parameter.

   ```bash
   curl http://localhost:8080/status?id=some-unique-job-id
   ```

   You will receive a response with the job status and result:

   ```json
   {
     "id": "some-unique-job-id",
     "status": "Completed",
     "result": "Processed data for job some-unique-job-id"
   }
   ```

## Explanation

This example demonstrates a job queue system where jobs are processed asynchronously using a pool of workers. When a job is submitted, it is added to a Redis-backed queue. Workers continuously poll the queue for new jobs, process them, and update their status in Redis. This allows the main application to remain responsive while handling long-running tasks in the background.

Will post full blog article later. In the meantime if you have any question you could raise them as issues on this repository or send me an email on [ahmad.alkadri@outlook.com](mailto:ahmad.alkadri@outlook.com).
