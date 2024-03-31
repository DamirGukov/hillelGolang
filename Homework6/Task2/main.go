package main

import (
	"flag"
	"fmt"
	"sync"
	"time"
)

func main() {
	numReceivers := flag.Int("n", 3, "number of receivers")
	flag.Parse()

	messageStream := make(chan string)
	var wg sync.WaitGroup

	numMessages := 10

	go func() {

		for i := 0; i <= numMessages; i++ {
			message := fmt.Sprintf("Message %d", i)
			messageStream <- message
			time.Sleep(100 * time.Millisecond)
		}
		close(messageStream)
	}()

	for r := 1; r <= *numReceivers; r++ {
		wg.Add(1)
		go func(receiverID int) {
			defer wg.Done()
			for msg := range messageStream {
				fmt.Printf("Receiver %d received: %s\n", receiverID, msg)
			}
		}(r)
	}

	wg.Wait()
}
