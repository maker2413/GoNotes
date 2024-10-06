package main

import (
	"fmt"
	"io"
	"os"
)

// By using io.Writer we are still able to use this function in main by passing in os.Stdout.
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func main() {
	Greet(os.Stdout, "Ethan")
}
