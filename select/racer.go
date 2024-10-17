package main

import (
	"fmt"
	"net/http"
	"time"
)

var tenSecondTimeout = 10 * time.Second

func Racer(a, b string) (winner string, err error) {
	return ConfigurableRacer(a, b, tenSecondTimeout)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, err error) {
	// select allows you to wait on multiple channels. The first one to send a value
	// wins and the code underneath the case is executed.
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	// time.After will trigger at "timeout" seconds and "win" our select race unless one of our
	// other cases finishes first.
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

// ping will create a chan of type struct{} and return it.
// struct{} is used instead of another type like bool because chan struct{}
// is the smallest data type available from a memory perspective.
func ping(url string) chan struct{} {
	// We use make to create our channel because if we did var ch chan struct{} our channel will
	// block forever when we try to send (<-) ping to it. This is because var will initialize
	// with the "zero" value of the type. For instance this will be 0 for ints and "" for strings.
	// For channels the "zero" value is nil and nil can not have anything sent to it (<-).
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}
