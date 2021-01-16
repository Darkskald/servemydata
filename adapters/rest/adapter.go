package rest

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Adapter is the interaction entity for handling incoming http requests.
type Adapter struct {
	r *mux.Router
}

func NewAdapter() Adapter {
	return Adapter{mux.NewRouter()}
}

// ListenAndServe is the central entry point of the REST server.
func (a Adapter) ListenAndServe() {
	log.Printf("Listening on http://0.0.0.0%s\n", ":8080")
	_ = http.ListenAndServe(":8080", a.r)
}
