package main

import (
	"context"
	"fmt"
	"time"
)

func child(ctx context.Context, name string) {
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
	ctx, cancel := context.WithTimeout(context.Background(), 6*time.Second)
	defer cancel()

	go child(ctx, "A")
	go child(ctx, "B")

	<-ctx.Done()
	fmt.Println("Parent done")
}

func main() {
	parent()
}
