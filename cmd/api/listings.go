package main

import (
	"errors"
	"fmt"
	"letsgofurther/internal/data"
	"letsgofurther/internal/validator"
	"net/http"
	"strconv"
)

/* GET */
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

/* POST */
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

/* PATCH */
func (app *application) patchListingById(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil || id < 1 {
		app.notFoundResponse(w, r)
		return
	}

	// Fetch the existing listing record from the database, sending a 404 Not Found // response to the client if we couldn't find a matching record.
	listing, err := app.models.Listings.Select(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrNotFoundRecord):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	// If the request contains a X-Expected-Version header, verify that the listing
	// version in the database matches the expected version specified in the header.
	if r.Header.Get("X-Expected-Version") != "" {
		if strconv.FormatInt(int64(listing.Version), 32) != r.Header.Get("X-Expected-Version") {
			app.editConflictResponse(w, r)
			return
		}
	}

	var input struct {
		Title       *string  `json:"title"`
		Description *string  `json:"description"`
		Price       *int64   `json:"price"`
		Categories  []string `json:"categories"`
	}

	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	//if the pointer input.Title is nil means the argument is not empty
	if input.Title != nil {
		listing.Title = *input.Title
	}
	if input.Description != nil {
		listing.Description = *input.Description
	}
	if input.Price != nil {
		listing.Price = *input.Price
	}
	if input.Categories != nil {
		listing.Categories = input.Categories
	}

	v := validator.New()
	data.ValidateListing(v, listing)
	if !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	//save to db:
	err = app.models.Listings.Update(listing)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrEditConflict):
			app.editConflictResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusCreated, envelope{"listing": listing}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

/* DELETE */
func (app *application) deleteListingById(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil || id < 1 {
		app.notFoundResponse(w, r)
		return
	}

	err = app.models.Listings.Delete(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrNotFoundRecord):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusNoContent, envelope{"message": "deleted!"}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
