package main

import (
	"net/http"

	httpserver "github.com/maker2413/GoNotes/scalingAcceptance/adapters/httpserver"
)

func main() {
	handler := http.HandlerFunc(httpserver.Handler)
	http.ListenAndServe(":8080", handler)
}
