package render

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/jseow5177/bed_and_breakfast/internals/config"
	"github.com/jseow5177/bed_and_breakfast/internals/models"
	"github.com/justinas/nosurf"
)

var pathToTemplate = "./templates"

var functions = template.FuncMap{}

// An AppConfig accessible in the render package
var app *config.AppConfig

// RegisterAppConfig saves a reference to the AppConfig in main.go
func RegisterAppConfig(a *config.AppConfig) {
	app = a
}

// CreateDefaultData adds default data that is needed in templates
// For example, the CSRF token
func CreateDefaultData(td *models.TemplateData, r *http.Request) *models.TemplateData{
	td.Error = app.Session.PopString(r.Context(), "error") // Extract error message from session
	td.Flash = app.Session.PopString(r.Context(), "flash") // Extract flash (success) message from session
	td.Warning = app.Session.PopString(r.Context(), "warning") // Extract warning message from session
	td.CSRFToken = nosurf.Token(r)
	return td
}

// Template renders a html template.
func Template(w http.ResponseWriter, r *http.Request, tmpl string, td *models.TemplateData) error {

	// A map from template name to the template files
	var tc map[string]*template.Template

	if app.UseCache {
		tc = app.TemplateCache
	} else {
		tc, _ = CreateTemplateCache() // Used in development mode to live reload changes in templates
	}

	t, ok := tc[tmpl] // Get the template of requested page

	td = CreateDefaultData(td, r)

	if !ok {
		return errors.New("Could not get template from template cache")
	}

	err := t.Execute(w, td) // Write into ResponseWriter

	if err != nil {
		fmt.Println("Error returning template")
		return err
	}

	return nil
}

// CreateTemplateCache creates a template cache as a map
func CreateTemplateCache()(map[string]*template.Template, error) {

	cache := make(map[string]*template.Template)

	// Glob returns the names of all files matching pattern or nil if there is no matching file.
	// Glob ignores file system errors such as I/O errors reading directories.
	// The only possible returned error is ErrBadPattern, where pattern is malformed.
	pages, err := filepath.Glob(fmt.Sprintf("%s/*.page.html", pathToTemplate))
	if err != nil {
		return cache, err
	}

	for _, page := range pages {
		// Base returns the last element of path.
		// Trailing path separators are removed before extracting the last element.
		filename := filepath.Base(page)

		// New allocates a new empty HTML template with the given name.
		newTemplate := template.New(filename)

		// Funcs adds the elements of the argument map to the template's function map.
		// A function map (FuncMap) is a map of functions that can be used in the template.
		ts, err := newTemplate.Funcs(functions).ParseFiles(fmt.Sprintf("%s/", pathToTemplate) + filename)
		if err != nil {
			return cache, err
		}

		// Check if there are any layouts
		layouts, err := filepath.Glob(fmt.Sprintf("%s/layouts/*.layout.html", pathToTemplate))
		if len(layouts) > 0 {
			// (*t Template) ParseGlob(pattern string) parses the template definitions in the files
			// identified by pattern and associates the resulting templates with t.
			ts, err = ts.ParseGlob(fmt.Sprintf("%s/layouts/*.layout.html", pathToTemplate))
			if err != nil {
				return cache, err
			}
		}

		// Save template into cache
		cache[filename] = ts
	}

	return cache, nil
}