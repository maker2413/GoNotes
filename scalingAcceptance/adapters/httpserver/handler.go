package httpserver

import (
	"fmt"
	"net/http"

	scaling_acceptance "github.com/maker2413/GoNotes/scalingAcceptance"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	fmt.Fprint(w, scaling_acceptance.Greet(name))
}
