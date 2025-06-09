package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Producer function produces numbers and sends them to channel
func producer(id int, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for range 5 {
		num := rand.Intn(100)
		fmt.Printf("Producer %d produced: %d\n", id, num)
		ch <- num
		time.Sleep(time.Millisecond * 500) // simulate work
	}
}

// Consumer function receives numbers from channel and processes them
func consumer(id int, ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range ch { // range on channel reads until channel is closed
		fmt.Printf("Consumer %d consumed: %d\n", id, num)
		time.Sleep(time.Millisecond * 800) // simulate processing
	}
}

func main() {

	ch := make(chan int, 10) // buffered channel with capacity 10

	var wg sync.WaitGroup

	// Start 3 producers
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go producer(i, ch, &wg)
	}

	// Start 2 consumers
	for i := 1; i <= 2; i++ {
		wg.Add(1)
		go consumer(i, ch, &wg)
	}

	// Wait for all producers to finish producing
	go func() {
		wg.Wait()  // Wait for producers and consumers (but consumers block on channel)
		close(ch)  // Close channel to signal consumers no more data will come
	}()

	// Wait a bit before exit (or block on something else)
	time.Sleep(7 * time.Second)
	fmt.Println("Main: Finished")
}
