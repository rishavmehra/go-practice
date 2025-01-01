package main

import "fmt"

func main() {
	// fmt.Printf("%v", result)
	// nilChannel()
	// simpleChannel()
	// unodirectionalchannel()
	// bufferchannelwithoutloop()
	// bufferchannelwithloop()
	nilChannel()

}

// func bufferchannelwithloop() {
// 	ch := make(chan int, 2)

// 	go func(ch chan int) {
// 		ch <- 21
// 		ch <- 22

// 		defer close(ch)
// 	}(ch)

// 	for value := range ch {
// 		fmt.Println(value)
// 	}

// }

// func bufferchannelwithoutloop() {
// 	var ch = make(chan int, 2)

// 	ch <- 21
// 	ch <- 22

// 	fmt.Println(<-ch)
// 	fmt.Println(<-ch)
// }

// func unidirectionalchannel() {
// 	var ch = make(chan int)

// 	go func(ch chan<- int) {
// 		ch <- 1
// 	}(ch)

// 	fmt.Println(<-ch)

// }

func nilChannel() {
	var ch = make(chan int)

	// go func() {
	ch <- 1
	// close(ch)
	// }()

	fmt.Println(<-ch)
}

// func simpleChannel() {
// 	var ch = make(chan int)

// 	go func() {
// 		ch <- 1
// 	}()

// 	// NOTE: if you comment this line. You will not be able to receive the result but code will not crash
// 	fmt.Println(<-ch)
// }
