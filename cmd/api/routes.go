package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router {
	router := httprouter.New()
	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodGet, "/v1/listings", app.getAllListings)
	router.HandlerFunc(http.MethodPost, "/v1/listings", app.postListing)
	router.HandlerFunc(http.MethodGet, "/v1/listings/:id", app.getListingById)
	router.HandlerFunc(http.MethodPatch, "/v1/listings/:id", app.patchListingById)
	router.HandlerFunc(http.MethodDelete, "/v1/listings/:id", app.deleteListingById)

	return router
}
