package main

import (
	"fmt"
	"time"
)

func listenToChan(ch chan int) {
	for {
		// print a got data message
		i := <-ch
		fmt.Println("Got", i, "from channel")

		// simulate doing a lot of work
		time.Sleep(1 * time.Second)
	}
}

func main() {
	// This buffered channel is allowed to have 10 things in it.
	ch := make(chan int, 10)

	go listenToChan(ch)

	for i := 0; i <= 100; i++ {
		// The first 10 times through this loop, things go quickly; after that,
		// things slow down.
		fmt.Println("sending", i, "to channel...")
		ch <- i
		fmt.Println("sent", i, "to channel!")
	}

	fmt.Println("Done!")
	close(ch)
}
