package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	slowServer := makeDelayedServer(20 * time.Millisecond)
	fastServer := makeDelayedServer(0 * time.Millisecond)

	// defer will call execute at the end of the containing functions execution
	defer slowServer.Close()
	defer fastServer.Close()

	slowURL := slowServer.URL
	fastURL := fastServer.URL

	want := fastURL
	got := Racer(slowURL, fastURL)

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	// httptest.NewServer takes an http.HandlerFunc which we are sending via an anonymous function
	// http.HandlerFunc is a type that looks like:
	// type HandlerFunc func(ResponseWriter, *Request).
	// This fake testing webserver will have a url of: http://127.0.0.1:<random port>
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		// WriteHeader is used to return an OK status (200)
		w.WriteHeader(http.StatusOK)
	}))
}
