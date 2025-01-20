package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	fmt.Println("Basic Calculator")
	fmt.Println("----------------")

	doneChan := make(chan bool)
	defer close(doneChan)

	go readUserInput(os.Stdin, doneChan)
	<-doneChan
}

func readUserInput(input io.Reader, doneChan chan bool) {
	scanner := bufio.NewScanner(input)

	for {
		fmt.Println("(type q to quit)")
		fmt.Print("-> ")
		response, done := checkNumbers(scanner)

		if done {
			doneChan <- true
			return
		}

		fmt.Println(response)
	}
}

func checkNumbers(scanner *bufio.Scanner) (string, bool) {
	scanner.Scan()

	if strings.EqualFold(scanner.Text(), "q") {
		return "", true
	}

	return scanner.Text(), false
}
