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
	render.Template(w, "home.page.html", &models.TemplateData{})
}

// About handles request to About page
func (*Repository)About(w http.ResponseWriter, r *http.Request) {
	render.Template(w, "about.page.html", &models.TemplateData{})
}

// Generals handles request to the General's Quarter room page
func (*Repository)Generals(w http.ResponseWriter, r *http.Request) {
	render.Template(w, "generals.page.html", &models.TemplateData{})
}

// Majors handles request to the Major's Suite room page
func (*Repository)Majors(w http.ResponseWriter, r *http.Request) {
	render.Template(w, "majors.page.html", &models.TemplateData{})
}

// Availability handles request to the page to search for room availability
func (*Repository)Availability(w http.ResponseWriter, r *http.Request) {
	render.Template(w, "search-availability.page.html", &models.TemplateData{})
}

// Contact handles request to the contact page
func (*Repository)Contact(w http.ResponseWriter, r *http.Request) {
	render.Template(w, "contact.page.html", &models.TemplateData{})
}