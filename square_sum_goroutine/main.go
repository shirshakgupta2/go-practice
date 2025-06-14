package main

import (
	"fmt"
	"sync"
)

// var mutex sync.Mutex
// func worker(ch chan int,wg *sync.WaitGroup,sum *int){
// 	defer wg.Done()
// 	for i := range ch {
// 		square := i * i
// 		mutex.Lock()
// 		*sum += square
// 		mutex.Unlock()
// 	}
// }

// func main() {
// 	var wg sync.WaitGroup
// 	ch := make(chan int)
// 	var sum int
// 	for i:= 1;i<=4;i++{
// 		wg.Add(1)
// 		go worker(ch,&wg,&sum)
// 	}
// 	for i:= 1;i<=100;i++{
// 		ch<-i
// 	}
// 	close(ch)
// 	wg.Wait()
// 	fmt.Println("Sum:",sum)

// }




func worker(ch <-chan int, wg *sync.WaitGroup, result chan<- int) {
	defer wg.Done()
	localSum := 0
	for i := range ch {
		localSum += i * i
		fmt.Println("localsum:",localSum)
	}
	fmt.Println("reading channel finished")
	result <- localSum
}

func main() {
	const workers = 4
	ch := make(chan int)
	result := make(chan int, workers)
	var wg sync.WaitGroup

	// Start workers
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go worker(ch, &wg, result)
	}

	// Send data
	for i := 1; i <= 100; i++ {
		ch <- i
	}
	close(ch)

	// Wait for workers to finish
	wg.Wait()
	close(result)

	// Aggregate results
	totalSum := 0
	for partial := range result {
		totalSum += partial
	}

	fmt.Println("Sum:", totalSum)
}
