package main

import (
	"servemydata/adapters/inmem"
	"servemydata/adapters/rest"
	"servemydata/usecases"
)

func main() {

	// set up use cases and repo
	repo := inmem.NewInMemoryRepo()
	adder := usecases.NewAddSpectrumService(&repo)

	// set up REST adapter
	adp := rest.NewAdapter()
	adp.HandleFunc("/", adp.MakeIndexHandler())

	adp.HandleFunc("/spectrum", adp.MakeAddSpectrumHandler(adder)).Methods("POST")
	// todo get spectrum handling
	// todo list spectra handling
	// todo delete spectra handling

	adp.ListenAndServe()

}
