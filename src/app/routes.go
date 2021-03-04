package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/jseow5177/bed_and_breakfast/internals/handlers"
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
	mux.Get("/generals-quarter", handlers.Repo.Generals)

	mux.Get("/make-reservation", handlers.Repo.Reservation)
	mux.Post("/make-reservation", handlers.Repo.PostReservation)
	mux.Get("/reservation-summary", handlers.Repo.ReservationSummary)

	mux.Get("/search-availability", handlers.Repo.Availability)
	mux.Post("/search-availability", handlers.Repo.PostAvailability)
	mux.Post("/json-test", handlers.Repo.AvailabilityJSON)

	// Initiates FileServer Handler
	// http.Dir() creates a FileSystem that points to the application static files
	fileServer := http.FileServer(http.Dir("./static/"))

	// http.StripPrefix removes "/static/" from the request URL
	mux.Handle("/static/*", http.StripPrefix("/static/", fileServer))

	return mux
}