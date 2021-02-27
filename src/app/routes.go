package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/jseow5177/bed_and_breakfast/pkg/handlers"
)


func routes() http.Handler {
	// chi routing
	mux := chi.NewRouter() // Create chi multiplexer

	// A middleware that recovers from panics, logs the panic and returns a HTTP500 status if possible
	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	return mux
}