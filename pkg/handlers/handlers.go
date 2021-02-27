package handlers

import (
	"net/http"

	"github.com/jseow5177/bed_and_breakfast/pkg/config"
	"github.com/jseow5177/bed_and_breakfast/pkg/models"
	"github.com/jseow5177/bed_and_breakfast/pkg/render"
)

// Repository type
type Repository struct {
	App *config.AppConfig
}

// Repo the repository used by the handlers
var Repo *Repository

// CreateNewRepo creates a new repository
func CreateNewRepo(a *config.AppConfig) *Repository {
	return &Repository {
		App: a,
	}
}

// RegisterRepo set the repository for the handlers
func RegisterRepo(r *Repository) {
	Repo = r
}

// Home handles request to root route
func (*Repository)Home(w http.ResponseWriter, r *http.Request) {
	remoteIP := r.RemoteAddr // Get IP Address of request
	Repo.App.Session.Put(r.Context(), "remote_ip", remoteIP) // Store IP Address into session

	render.Template(w, "home.page.html", &models.TemplateData{})
}

// About handles request to About page
func (*Repository)About(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["greet"] = "Hello World"

	remoteIP := Repo.App.Session.GetString(r.Context(), "remote_ip")
	stringMap["remote_ip"] = remoteIP

	render.Template(w, "about.page.html", &models.TemplateData{
		StringMap: stringMap,
	})
}