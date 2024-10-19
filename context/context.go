package main

import (
	"context"
	"fmt"
	"net/http"
)

type Store interface {
	Fetch(ctx context.Context) (string, error)
}

// The function Server takes a Store and returns us a http.HandlerFunc.
func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := store.Fetch(r.Context())

		if err != nil {
			return // todo: log error however you like
		}

		fmt.Fprint(w, data)
	}
}
