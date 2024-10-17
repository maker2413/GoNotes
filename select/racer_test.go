package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("compare speeds of servers, returning the url of the fastest one", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(0 * time.Millisecond)

		// defer will call execute at the end of the containing functions execution
		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got, err := Racer(slowURL, fastURL)

		if err != nil {
			t.Fatalf("did not expect an error but got one %v", err)
		}

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("returns an error if a server doesn't respond within the specified time", func(t *testing.T) {
		server := makeDelayedServer(25 * time.Millisecond)

		defer server.Close()

		_, err := ConfigurableRacer(server.URL, server.URL, 20*time.Millisecond)

		if err == nil {
			t.Errorf("expected an error, but didn't get one")
		}
	})
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
