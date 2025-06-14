package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func worker(ctx context.Context, id int, jobs <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Worker", id, "canceled:", ctx.Err())
			return
		case job, ok := <-jobs:
			if !ok {
				return
			}
			fmt.Printf("Worker %d processing job %d\n", id, job)
			time.Sleep(time.Millisecond * 300) // simulate work
		}
	}
}


func main() {
	jobs := make(chan int, 10)
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go worker(ctx, i, jobs, &wg)
	}

	for j := 1; j <= 10; j++ {
		jobs <- j
	}
	close(jobs)

	wg.Wait()
}
