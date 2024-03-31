package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	dataChannel := make(chan int, 100)
	var wg sync.WaitGroup
	semaphore := make(chan struct{}, 10)

	for i := 0; i < 100; i++ {
		go func(id int) {
			for j := 0; j < 10; j++ {
				semaphore <- struct{}{}
				wg.Add(1)
				go func(id, j int) {
					defer wg.Done()
					defer func() { <-semaphore }()
					sleepDuration := time.Duration(rand.Intn(991)+10) * time.Millisecond
					time.Sleep(sleepDuration)
					dataChannel <- id*100 + j
				}(id, j)
			}

			fmt.Printf("Goroutine %d finished\n", id)
		}(i)
	}

	go func() {
		for value := range dataChannel {
			fmt.Printf("Received value: %d\n", value)
		}
	}()

	wg.Wait()
	close(dataChannel)
}
