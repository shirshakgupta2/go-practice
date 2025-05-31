package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Job represents a unit of work
type Job struct {
	ID       int
	Workload int // simulated duration
}

// Worker simulates a worker that processes jobs
func worker(id int, jobs <-chan Job, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		fmt.Printf("[Worker %d] Started Job %d (Workload: %dms)\n", id, job.ID, job.Workload)
		time.Sleep(time.Duration(job.Workload) * time.Millisecond) // simulate processing
		fmt.Printf("[Worker %d] Finished Job %d\n", id, job.ID)
	}
}

func main() {
	const (
		numWorkers = 3
		numJobs    = 10
	)

	jobs := make(chan Job, numJobs)
	var wg sync.WaitGroup

	// Start the worker pool
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, jobs, &wg)
	}

	// Send jobs
	for j := 1; j <= numJobs; j++ {
		workload := rand.Intn(1000) + 500 // 500ms to 1500ms
		jobs <- Job{ID: j, Workload: workload}
	}
	close(jobs) // no more jobs to send

	wg.Wait()
	fmt.Println("All jobs completed.")
}
