package rest

import (
	"fmt"
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
	log.Fatal(http.ListenAndServe(":8080", a.r))
}

// HandleFunc is a convenience function for the registration of new handler functions.
func (a Adapter) HandleFunc(path string, f func(w http.ResponseWriter, r *http.Request)) *mux.Route {
	return a.r.NewRoute().Path(path).HandlerFunc(f)
}

func (a Adapter) MakeIndexHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		msg := "<h1>Welcome to servemydata </h1>"
		fmt.Fprintf(w, msg)
	}
}
