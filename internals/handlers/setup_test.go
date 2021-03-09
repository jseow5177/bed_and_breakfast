package handlers

import (
	"encoding/gob"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/jseow5177/bed_and_breakfast/internals/config"
	"github.com/jseow5177/bed_and_breakfast/internals/models"
	"github.com/jseow5177/bed_and_breakfast/internals/render"
)

var app config.AppConfig
var functions = template.FuncMap{}

func getRoutes() http.Handler {
	gob.Register(models.Reservation{})

	app.InProduction = false // Change to true when in production

	// Initialise new session manager
	session := scs.New()
	session.Lifetime = 24 * time.Hour // Session lasts for 24 hours
	session.Cookie.Persist = true // Session persists when user closes window or browser
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction // No HTTPS when in development

	app.Session = session // Stores SessionManager in AppConfig so that it is accessible in other packages

	tc, err := CreateTestTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	app.TemplateCache = tc // Store the template cache
	app.UseCache = true

	// Gives the render package access to the AppConfig instance
	render.RegisterAppConfig(&app)

	// Create a new repository for the handler package
	repo := CreateNewRepo(&app)
	// Gives the handlers package access to the newly created repo, which holds the AppConfig instance
	RegisterRepo(repo)

	// chi routing
	mux := chi.NewRouter() // Create chi multiplexer

	// A middleware that recovers from panics, logs the panic and returns a HTTP500 status if possible
	mux.Use(middleware.Recoverer)
	mux.Use(SessionLoad)

	mux.Get("/", Repo.Home)
	mux.Get("/about", Repo.About)
	mux.Get("/generals-quarter", Repo.Generals)

	mux.Get("/make-reservation", Repo.Reservation)
	mux.Post("/make-reservation", Repo.PostReservation)
	mux.Get("/reservation-summary", Repo.ReservationSummary)

	mux.Get("/search-availability", Repo.Availability)
	mux.Post("/search-availability", Repo.PostAvailability)
	mux.Post("/json-test", Repo.AvailabilityJSON)

	// Initiates FileServer Handler
	// http.Dir() creates a FileSystem that points to the application static files
	fileServer := http.FileServer(http.Dir("./static/"))

	// http.StripPrefix removes "/static/" from the request URL
	mux.Handle("/static/*", http.StripPrefix("/static/", fileServer))

	return mux
}

// SessionLoad is a middleware that automatically loads and saves session
// data for the current request, and communicates the session token to and from
// the client in a cookie.
func SessionLoad(next http.Handler) http.Handler {
	return app.Session.LoadAndSave(next)
}

// CreateTestTemplateCache creates a template cache as a map
func CreateTestTemplateCache()(map[string]*template.Template, error) {

	cache := make(map[string]*template.Template)

	// Glob returns the names of all files matching pattern or nil if there is no matching file.
	// Glob ignores file system errors such as I/O errors reading directories.
	// The only possible returned error is ErrBadPattern, where pattern is malformed.
	pages, err := filepath.Glob("../../templates/*.page.html")
	if err != nil {
		return cache, err
	}

	for _, page := range pages {
		// Base returns the last element of path.
		// Trailing path separators are removed before extracting the last element.
		filename := filepath.Base(page)

		// New allocates a new empty HTML template with the given name.
		// template.Name() returns the name of the template.
		newTemplate := template.New(filename)

		// Funcs adds the elements of the argument map to the template's function map.
		// A function map (FuncMap) is a map of functions that can be used in the template.
		ts, err := newTemplate.Funcs(functions).ParseFiles("../../templates/" + filename)
		if err != nil {
			return cache, err
		}

		// Check if there are any layouts
		layouts, err := filepath.Glob("../../templates/layouts/*.layout.html")
		if len(layouts) > 0 {
			// (*t Template) ParseGlob(pattern string) parses the template definitions in the files
			// identified by pattern and associates the resulting templates with t.
			ts, err = ts.ParseGlob("../../templates/layouts/*.layout.html")
			if err != nil {
				return cache, err
			}
		}

		// Save template into cache
		cache[filename] = ts
	}

	return cache, nil
}