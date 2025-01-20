package main

import (
	"bytes"
	"testing"
)

func Test_readUserInput(t *testing.T) {
	doneChan := make(chan bool)
	defer close(doneChan)

	var stdin bytes.Buffer

	go readUserInput(&stdin, doneChan)
	<-doneChan
}
