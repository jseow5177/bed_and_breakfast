package handlers

import (
	"encoding/gob"
	"html/template"
	"log"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"testing"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi"
	"github.com/jseow5177/bed_and_breakfast/internals/config"
	"github.com/jseow5177/bed_and_breakfast/internals/models"
	"github.com/jseow5177/bed_and_breakfast/internals/render"
)

var app config.AppConfig
var functions = template.FuncMap{}
var ts *httptest.Server

func TestMain(m *testing.M) {
	gob.Register(models.Reservation{})

	app.InProduction = false

	session := scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := CreateTestTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
	}

	app.TemplateCache = tc
	app.UseCache = true

	render.RegisterAppConfig(&app)

	repo := CreateNewRepo(&app)
	RegisterRepo(repo)

	mux := chi.NewRouter()

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

	fileServer := http.FileServer(http.Dir("./static/"))

	mux.Handle("/static/*", http.StripPrefix("/static/", fileServer))

	ts = httptest.NewServer(mux)
}

func SessionLoad(next http.Handler) http.Handler {
	return app.Session.LoadAndSave(next)
}

func CreateTestTemplateCache()(map[string]*template.Template, error) {

	cache := make(map[string]*template.Template)

	pages, err := filepath.Glob("../../templates/*.page.html")
	if err != nil {
		return cache, err
	}

	for _, page := range pages {
		filename := filepath.Base(page)

		newTemplate := template.New(filename)

		ts, err := newTemplate.Funcs(functions).ParseFiles("../../templates/" + filename)
		if err != nil {
			return cache, err
		}

		layouts, err := filepath.Glob("../../templates/layouts/*.layout.html")
		if len(layouts) > 0 {
			ts, err = ts.ParseGlob("../../templates/layouts/*.layout.html")
			if err != nil {
				return cache, err
			}
		}

		cache[filename] = ts
	}

	return cache, nil
}