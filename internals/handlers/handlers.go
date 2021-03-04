package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/jseow5177/bed_and_breakfast/internals/config"
	"github.com/jseow5177/bed_and_breakfast/internals/forms"
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
	render.Template(w, r, "make-reservation.page.html", &models.TemplateData{
		Form: forms.New(nil), // Returns an empty form object
	})
}

// PostReservation handles the post of a reservation form
func (*Repository)PostReservation(w http.ResponseWriter, r *http.Request) {
	// ParseForm so that form values for available in r.Form and r.PostForm
	// Both r.Form and r.PostForm are of type url.Values
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		return
	}

	form := forms.New(r.PostForm)
	form.Required("first_name", "last_name", "email")
	form.MinLength("first_name", 3)
	form.IsEmail("email")

	if !form.Valid() {
		render.Template(w, r, "make-reservation.page.html", &models.TemplateData{
			Form: form, // Returns form with errors object and values submitted by user
		})
		return
	}

	reservation := models.Reservation{
		FirstName: form.Values.Get("first_name"),
		LastName: form.Values.Get("last_name"),
		Email: form.Values.Get("email"),
		Phone: form.Values.Get("phone"),
	}

	// Store data into session
	Repo.App.Session.Put(r.Context(), "reservation", reservation)
	http.Redirect(w, r, "/reservation-summary", http.StatusSeeOther) // 303 redirect
}

// ReservationSummary displays the summary of successful reservation made by user
func (*Repository)ReservationSummary(w http.ResponseWriter, r *http.Request) {
	// Session.Pop returns a value of type interface{}
	// It needs to be type asserted
	// A type assertion provides access to an interface value's underlying concrete value
	// t := i.(T)
	// The statement above asserts that the interface value i holds the concrete type T and assigns the underlying
	// T value to the variable t.
	reservation, ok := Repo.App.Session.Pop(r.Context(), "reservation").(models.Reservation)
	if !ok {
		// Pass error message to session
		Repo.App.Session.Put(r.Context(), "error", "You have not made a reservation!")
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	data := make(map[string]interface{})
	data["reservation"] = reservation

	render.Template(w, r, "reservation-summary.page.html", &models.TemplateData{
		Data: data,
	})
}