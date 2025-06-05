package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func child(ctx context.Context, name string, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Child", name, "stopped:", ctx.Err())
			return
		default:
			fmt.Println("Child", name, "working...")
			time.Sleep(300 * time.Millisecond)
		}
	}
}

func parent() { 
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	var wg sync.WaitGroup
	wg.Add(2)
	go child(ctx, "A", &wg)
	go child(ctx, "B", &wg)

	<-ctx.Done()
	wg.Wait()
	fmt.Println("Parent context done:", ctx.Err())
	fmt.Println("Parent done")
}

func main() {
	parent()
}
