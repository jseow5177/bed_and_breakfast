package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/jseow5177/bed_and_breakfast/internals/config"
	"github.com/jseow5177/bed_and_breakfast/internals/handlers"
	"github.com/jseow5177/bed_and_breakfast/internals/models"
	"github.com/jseow5177/bed_and_breakfast/internals/render"
)

// Declare an AppConfig
var app config.AppConfig

const portNumber = ":8080"

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
	
	// Create a custom server
	srv := &http.Server{
		Addr: portNumber, // Port number
		Handler: routes(), // Register multiplexer
	}

	fmt.Println("App is listening at port", portNumber)
	
	// ListenAndServe always return a non-nil error
	// In other words, it will only return when it fails
	err = srv.ListenAndServe()
	log.Fatal(err)
}

func run() error {
	gob.Register(models.Reservation{})

	app.InProduction = false // Change to true when in production

	// Initialise new session manager
	session := scs.New()
	session.Lifetime = 24 * time.Hour // Session lasts for 24 hours
	session.Cookie.Persist = true // Session persists when user closes window or browser
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction // No HTTPS when in development

	app.Session = session // Stores SessionManager in AppConfig so that it is accessible in other packages

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot create template cache")
		return err
	}

	app.TemplateCache = tc // Store the template cache
	app.UseCache = false // Reload changes in HTML template without recompiling (use in development mode)

	// Gives the render package access to the AppConfig instance
	render.RegisterAppConfig(&app)

	// Create a new repository for the handler package
	repo := handlers.CreateNewRepo(&app)
	// Gives the handlers package access to the newly created repo, which holds the AppConfig instance
	handlers.RegisterRepo(repo)

	return nil
}