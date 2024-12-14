package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

// Can't be called TestMain, because that is already used by testing
func Test_main(t *testing.T) {
	stdOut := os.Stdout
	r, w, _ := os.Pipe()

	os.Stdout = w

	main()

	_ = w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdOut

	if !strings.Contains(output, "$34320.00") {
		t.Error("Wrong balanced retruned")
	}
}
