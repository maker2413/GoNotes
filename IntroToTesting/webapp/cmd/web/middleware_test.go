package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_application_addIPToContext(t *testing.T) {
	tests := []struct {
		headerName  string
		headerValue string
		addr        string
		emptyAddr   bool
	}{
		{"", "", "", false},
		{"", "", "", true},
		{"X-Forwarded-For", "192.3.2.1", "", false},
		{"", "", "hello:world", false},
	}

	// Create a dumby handler that we'll use to check the context
	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Make sure that the value exists in the context
		val := r.Context().Value(contextUserKey)
		if val == nil {
			t.Error(contextUserKey, "not present")
		}

		// Make sure we got a string back
		ip, ok := val.(string)
		if !ok {
			t.Error("not string")
		}

		t.Log(ip)
	})

	for _, e := range tests {
		// Create the handler to test
		handlerToTest := app.addIPToContext(nextHandler)

		req := httptest.NewRequest("GET", "http://testing", nil)

		if e.emptyAddr {
			req.RemoteAddr = ""
		}

		if len(e.headerName) > 0 {
			req.Header.Add(e.headerName, e.headerValue)
		}

		if len(e.addr) > 0 {
			req.RemoteAddr = e.addr
		}

		handlerToTest.ServeHTTP(httptest.NewRecorder(), req)
	}
}

func Test_application_ipFromContext(t *testing.T) {
	want := "whatever"

	// Get a context
	ctx := context.Background()

	// Put something in the context
	ctx = context.WithValue(ctx, contextUserKey, want)

	// Call the function
	got := app.ipFromContext(ctx)

	// Perform the test
	if got != want {
		t.Errorf("got: %q, expected: %q", got, want)
	}
}
