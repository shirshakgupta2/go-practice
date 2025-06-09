// package main

// import (
// 	"fmt"
// 	"time"
// )

// func worker(channel chan string, i int) {
// 	data := fmt.Sprint("worker complete", i)
// 	channel <- data
// }

// func main() {
// 	channel := make(chan string)
// 	for i := range 5 {

// 		go worker(channel, i)
// 	}

// 	for {
// 		select {
// 		case <-time.After(5 * time.Second):
// 			close(channel)
// 			return
// 		case msg := <-channel:
// 			fmt.Println(msg)
// 		}
// 	}
// }

package main

import (
	"fmt"
	"sync"
)

func worker(channel chan string, i int, wg *sync.WaitGroup) {
	defer wg.Done() // Ensure that we signal completion of the worker
	data := fmt.Sprint("worker complete", i)
	channel <- data
}

func main() {

	channel := make(chan string)
	var wg sync.WaitGroup
	go func() {
		for msg := range channel {
			fmt.Println(msg)
		}
	}()
	for i := range 5 {
		wg.Add(1)
		go worker(channel, i, &wg)
	}

	wg.Wait()
	close(channel)
	fmt.Println("All workers completed, channel closed.")
}
