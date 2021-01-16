package main

import "servemydata/adapters/rest"

func main() {

	adp := rest.NewAdapter()
	adp.ListenAndServe()
}
