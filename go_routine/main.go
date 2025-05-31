package main

import (
	"fmt"
	"time"
)

func main() {
	testchan := make(chan string)
	go worker(testchan)
	for {
		select {
		case <-time.After(5 * time.Second):
			fmt.Print("Timeout reached, exiting...\n")
			close(testchan)
			return
		case msg := <-testchan:
			fmt.Println(msg)
		}
	}
}

func worker(testchan chan string) {
	for i := 0; i < 10; i++ {
		testchan <- fmt.Sprintf("Worker: %d", i)
	}
}
