package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	if len(os.Args) <= 1 {
		log.Fatal("Please provide a valid file name(s) to read from!")
	}

	fmt.Printf("Opening: %s\n", os.Args[1:])
	for _, arg := range os.Args[1:] {
		f, err := os.Open(arg)

		if err != nil {
			log.Fatalf("File: %s does not exist!", arg)
		}

		io.Copy(os.Stdout, f)
	}
}
