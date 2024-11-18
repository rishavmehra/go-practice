package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	SingleProducerConsumer()

}

func SingleProducerConsumer() {
	ch := make(chan string)
	var wg sync.WaitGroup

	wg.Add(1)
	go SingleProducer(ch, &wg)
	wg.Add(1)
	go SingleConsumer(ch, &wg)
	wg.Wait()
	close(ch)
}

func SingleProducer(RiCh chan string, wg *sync.WaitGroup) {
	delay(8)
	RiCh <- "this the producer Message"
	wg.Done()
}

func SingleConsumer(ch chan string, wg *sync.WaitGroup) {
	select {
	case msg := <-ch:
		delay(1)
		fmt.Println("this is the Consumer, this will receive the Producer msg and here is the Producer msg -> ", msg)
	case <-time.After(10 * time.Second):
		fmt.Printf("stop listening to the channel after %d milli seconds", 1)
		wg.Done()
	}
	fmt.Println("Consumer finishes the job")
	wg.Done()
}

func delay(delay int) {
	time.Sleep(time.Duration(delay) * time.Second)
}
