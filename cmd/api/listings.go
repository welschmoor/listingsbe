package main

import (
	"fmt"
	"net/http"
)

func (app *application) getListingById(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "show the details of movie %d\n", id)
}

func (app *application) postListing(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new movie")
}
