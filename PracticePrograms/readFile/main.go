package main

import (
	"fmt"
	"io"
	"os"
)

const fileError string = "Please provide a valid file name to read from!"

func main() {
	if len(os.Args) <= 1 {
		fmt.Println(fileError)
		os.Exit(1)
	}

	fmt.Printf("Opening: %s\n", os.Args[1])
	f, err := os.Open(os.Args[1])

	if err != nil {
		fmt.Println(fileError)
		os.Exit(1)
	}

	io.Copy(os.Stdout, f)
}
