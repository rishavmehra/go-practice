package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	MultiProducerConsumer(100, 50)

}

func MultiProducerConsumer(producerSize int, consumerSize int) {
	ch := make(chan string)
	var wg sync.WaitGroup

	for i := 0; i < producerSize; i++ {
		wg.Add(1)
		go Producer(i, ch, &wg)
	}
	for i := 0; i < producerSize; i++ {
		wg.Add(1)
		go Consumer(i, ch, &wg)
	}

	wg.Wait()
	close(ch)
}

func Producer(index int, RiCh chan string, wg *sync.WaitGroup) {
	for i := 0; i < 5; i++ {
		RiCh <- fmt.Sprintf("Producer %v send %v", index, i)
	}
	wg.Done()
}

func Consumer(index int, ch chan string, wg *sync.WaitGroup) {
	done := false
	for !done {
		select {
		case msg, ok := <-ch:
			if !ok {
				done = true
			}
			fmt.Printf("Consumer %v Received: %s\n", index, msg)
		case <-time.After(10 * time.Second):
			fmt.Printf("stop listening to the channel after %d milli seconds", 1)
			wg.Done()
		}
	}
	wg.Done()
}
