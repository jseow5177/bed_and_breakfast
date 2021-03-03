package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/jseow5177/bed_and_breakfast/internals/config"
	"github.com/jseow5177/bed_and_breakfast/internals/models"
	"github.com/jseow5177/bed_and_breakfast/internals/render"
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
	render.Template(w, r, "home.page.html", &models.TemplateData{})
}

// About handles request to About page
func (*Repository)About(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "about.page.html", &models.TemplateData{})
}

// Generals handles request to the General's Quarter room page
func (*Repository)Generals(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "generals.page.html", &models.TemplateData{})
}

// Majors handles request to the Major's Suite room page
// func (*Repository)Majors(w http.ResponseWriter, r *http.Request) {
// 	render.Template(w, r, "majors.page.html", &models.TemplateData{})
// }

// Availability handles request to the page to search for room availability
func (*Repository)Availability(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "search-availability.page.html", &models.TemplateData{})
}

// PostAvailability handles post request from user to search for available rooms
func (*Repository)PostAvailability(w http.ResponseWriter, r *http.Request) {
	// By default, form data will be submitted as strings
	start := r.Form.Get("start_date")
	end := r.Form.Get("end_date")
	w.Write([]byte(fmt.Sprintf("Start %s, end %s", start, end)))
}

type jsonResponse struct {
	Ok bool
	Message string `json:"message"`
}

// AvailabilityJSON sample
func (*Repository)AvailabilityJSON(w http.ResponseWriter, r *http.Request) {
	resp := jsonResponse{
		Ok: true,
		Message: "Available!",
	}

	out, err := json.MarshalIndent(resp, "", "    ")
	if err != nil {
		log.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(out)
}

// Contact handles request to the contact page
func (*Repository)Contact(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "contact.page.html", &models.TemplateData{})
}

// Reservation handles request to the make-reservation page
func (*Repository)Reservation(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "make-reservation.page.html", &models.TemplateData{})
}