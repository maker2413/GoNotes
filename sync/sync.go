package main

import "sync"

type Counter struct {
	mu    sync.Mutex
	value int
}

// NewCounter returns a new Counter.
func NewCounter() *Counter {
	return &Counter{}
}

func (c *Counter) Inc() {
	// What this means is any goroutine calling Inc will acquire the lock on Counter if they
	// are first. All the other goroutines will have to wait for it to be Unlocked before
	// getting access.
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

func (c *Counter) Value() int {
	return c.value
}
