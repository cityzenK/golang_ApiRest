package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandle)
	router.HandlerFunc(http.MethodPost, "/v1/movies", app.createMovieHandle)
	router.HandlerFunc(http.MethodGet, "/v1/movies/:id", app.showMovieHandler)

	return router
}
