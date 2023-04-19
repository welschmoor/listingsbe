package main

import (
	"fmt"
	"letsgofurther/internal/data"
	"letsgofurther/internal/validator"
	"net/http"
	"time"
)

func (app *application) getListingById(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil || id < 1 {
		app.notFoundResponse(w, r)
		return
	}

	lis := data.Listing{
		ID:          123123,
		Title:       "Bike",
		Description: "Mint condition an all",
		Price:       120_00,
		Categories:  []string{"fahrrad", "freizeit"},
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"listing": lis}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) postListing(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title       string   `json:"title"`
		Description string   `json:"description"`
		Price       int64    `json:"price"`
		Categories  []string `json:"categories"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	v := validator.New()
	lis := data.Listing{
		Title:       input.Title,
		Description: input.Description,
		Price:       input.Price,
		Categories:  input.Categories,
	}
	data.ValidateListing(v, &lis)
	if !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}
	fmt.Fprintf(w, "%+v\n", input)
}
