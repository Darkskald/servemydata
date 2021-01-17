package rest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"servemydata/domain"
	"servemydata/usecases"

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

// ReadSpec2d is used to deserialize a JSON request body to a domain.Spectrum2d instance.
func (a Adapter) ReadSpec2d(r *http.Request) (domain.Spectrum2d, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return domain.Spectrum2d{}, err
	}

	var spectrum domain.Spectrum2d
	if err := json.Unmarshal(body, &spectrum); err != nil {
		return domain.Spectrum2d{}, err
	}

	return spectrum, nil
}

func (a Adapter) writeSpec2d(spec domain.Spectrum2d, w http.ResponseWriter) error {
	b, err := json.Marshal(spec)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	_, _ = w.Write(b)
	return nil
}

func (a Adapter) MakeIndexHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		msg := "<h1>Welcome to servemydata </h1>"
		fmt.Fprintf(w, msg)
	}
}

// usecase interaction

func (a Adapter) MakeAddSpectrumHandler(adder usecases.AddSpectrumService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		spectrum, err := a.ReadSpec2d(r)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		newSpec, err := adder.Run(spectrum)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if err = a.writeSpec2d(newSpec, w); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

	}
}
