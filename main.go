package main

import "servemydata/adapters/rest"

func main() {

	adp := rest.NewAdapter()
	adp.HandleFunc("/", adp.MakeIndexHandler())

	adp.ListenAndServe()

}
