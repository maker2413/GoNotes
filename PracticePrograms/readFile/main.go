package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

const fileError string = "Please provide a valid file name to read from!"

func main() {
	if len(os.Args) <= 1 {
		log.Fatal(fileError)
	}

	fmt.Printf("Opening: %s\n", os.Args[1])
	f, err := os.Open(os.Args[1])

	if err != nil {
		log.Fatal(fileError)
	}

	io.Copy(os.Stdout, f)
}
