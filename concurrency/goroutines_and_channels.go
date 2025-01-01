package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"runtime"
	"sync"
	"syscall"
	"time"
)

type CounterValue struct {
	ChuckID string
	Value   int
}

func main() {
	numCPUs := runtime.NumCPU()

	ctx, cancelCtx := context.WithCancel(context.Background())
	defer cancelCtx()

	dataChannel := make(chan *CounterValue)

	signalChannel := GetSignalChan()

	go func() {
		<-signalChannel
		fmt.Println("Terminating...")

		cancelCtx()
		close(signalChannel)
	}()

	wg := sync.WaitGroup{}
	wg.Add(numCPUs)

	for i := 0; i < numCPUs; i++ {
		go ChuckNorris(ctx, &wg, dataChannel, fmt.Sprintf("Chuck #%d", i), int32(i*10))
	}

	for {
		msg, open := <-dataChannel
		if !open {
			break
		}
		fmt.Printf("%s counts%d\n", msg.ChuckID, msg.Value)
	}

	wg.Wait()

}

func ChuckNorris(ctx context.Context, wg *sync.WaitGroup, dataChannel chan *CounterValue, id string, increment int32) {
	counter := int32(0)
	sent := true

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("%s has left the building\n", id)
			wg.Done()
			return
		default:

		}

		time.Sleep(2 * time.Second)

		if sent {
			counter += increment
		}

		select {
		case dataChannel <- &CounterValue{
			ChuckID: id,
			Value:   int(counter),
		}:
			sent = true
		default:
			sent = false
		}

	}
}

func GetSignalChan() chan os.Signal {
	signalChannel := make(chan os.Signal, 1)
	signal.Notify(signalChannel, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	return signalChannel
}
