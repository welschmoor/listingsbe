package main

import (
	"expvar"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()
	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)

	router.HandlerFunc(http.MethodGet, "/v1/listings", app.getAllListings)
	router.HandlerFunc(http.MethodPost, "/v1/listings", app.requirePermission("listings:write", app.postListing))
	router.HandlerFunc(http.MethodGet, "/v1/listings/:id", app.requirePermission("listings:read", app.getListingById))
	router.HandlerFunc(http.MethodPatch, "/v1/listings/:id", app.patchListingById)
	router.HandlerFunc(http.MethodDelete, "/v1/listings/:id", app.deleteListingById)

	router.HandlerFunc(http.MethodPost, "/v1/users", app.postUser)
	router.HandlerFunc(http.MethodPut, "/v1/users/activated", app.activateUserHandler)

	router.HandlerFunc(http.MethodPost, "/v1/tokens/activation", app.postActivationTokenHandler)
	router.HandlerFunc(http.MethodPost, "/v1/tokens/authentication", app.postAuthenticationTokenHandler)

	// Register a new GET /debug/vars endpoint pointing to the expvar handler.
	router.Handler(http.MethodGet, "/debug/vars", expvar.Handler())

	return app.metrics(app.recoverPanic(app.enableCORS(app.rateLimit(app.authenticate(router)))))
}
