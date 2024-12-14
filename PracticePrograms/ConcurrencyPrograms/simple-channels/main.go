package main

import (
	"fmt"
	"strings"
)

// shout has two parameters: a receive only chan ping, and a send only chan pong.
// Note the use of <- in function signature. It simply takes whatever string it
// gets from the ping channel, converts it to uppercase and appends a few
// exclamation marks, and then sends the transformed text to the pong channel.
func shout(ping <-chan string, pong chan<- string) {
	for {
		// Read from the ping channel. Note that the GoRoutine waits here -- it
		// blocks until something is received on this channel.
		s, ok := <-ping
		if !ok {
			// Do something.
		}

		pong <- fmt.Sprintf("%s!!!", strings.ToUpper(s))
	}
}

func main() {
	// Create two channels. Ping is what we send to, and pong is what comes back.
	ping := make(chan string)
	pong := make(chan string)

	// Start a goroutine
	go shout(ping, pong)

	fmt.Println("Type something and press ENTER (enter Q to quit)")

	for {
		// Print a prompt.
		fmt.Print("-> ")

		// Get user input.
		var userInput string
		_, _ = fmt.Scanln(&userInput)

		if strings.ToLower(userInput) == "q" {
			// Jump out of loop.
			break
		}

		ping <- userInput
		// Wait for a response.
		response := <-pong

		fmt.Println("Response:", response)
	}

	fmt.Println("All done. Closing channels.")
	close(ping)
	close(pong)
}
