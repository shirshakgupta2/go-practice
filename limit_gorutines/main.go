package main

import (
    "fmt"
    "runtime"
    "sync"
    "time"
)

func main() {
    const maxGoroutines = 4
    jobs := 10

    limitchan := make(chan struct{}, maxGoroutines)
    var wg sync.WaitGroup

    for i := 0; i < jobs; i++ {
        wg.Add(1)
        limitchan <- struct{}{} // acquire slot
        go func(jobID int) {
            defer wg.Done()
            defer func() { <-limitchan }() // release slot

            fmt.Printf("Job %d started | Running goroutines: %d\n", jobID, runtime.NumGoroutine())
            time.Sleep(2 * time.Second)
        }(i)
    }

    wg.Wait()
}
