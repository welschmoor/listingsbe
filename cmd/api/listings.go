package main

import (
	"errors"
	"fmt"
	"letsgofurther/internal/data"
	"letsgofurther/internal/validator"
	"net/http"
)

func (app *application) getListingById(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil || id < 1 {
		app.notFoundResponse(w, r)
		return
	}

	lis, err := app.models.Listings.Select(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrNotFoundRecord):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
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
	lis := &data.Listing{
		Title:       input.Title,
		Description: input.Description,
		Price:       input.Price,
		Categories:  input.Categories,
	}

	data.ValidateListing(v, lis)
	if !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	//save to db:
	err = app.models.Listings.Insert(lis)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	// When sending an HTTP response, we want to include a Location header to let the
	// client know which URL they can find the newly-created resource at. We make an
	// empty http.Header map and then use the Set() method to add a new Location header, // interpolating the system-generated ID for our new listing in the URL.
	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/listings/%d", lis.ID))

	// Write a JSON response with a 201 Created status code, the listing data in the // response body, and the Location header.
	err = app.writeJSON(w, http.StatusCreated, envelope{"listing": lis}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
