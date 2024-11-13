package main

import (
	"net/http"

	scaling_acceptance "github.com/maker2413/GoNotes/scalingAcceptance"
)

func main() {
	handler := http.HandlerFunc(scaling_acceptance.Handler)
	http.ListenAndServe(":8080", handler)
}
