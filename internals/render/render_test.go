package render

import (
	"net/http"
	"testing"

	"github.com/jseow5177/bed_and_breakfast/internals/models"
)


func TestCreateDefaultData(t *testing.T) {
	var td models.TemplateData

	r, err := getRequestWithSession()
	if err != nil {
		t.Fatal(err)
	}

	session.Put(r.Context(), "flash", "123")
	result := CreateDefaultData(&td, r)

	if result.Flash != "123" {
		t.Errorf("Expected flash value to be %s, but got %s", "123", result.Flash)
	}
}

func TestTemplate(t *testing.T) {
	r, err := getRequestWithSession()
	if err != nil {
		t.Fatal(err)
	}

	pathToTemplate = "../../templates"

	ww := mockWriter{}
	err = Template(&ww, r, "home.page.html", &models.TemplateData{})

	if err != nil {
		t.Error("Error in creating template for home.page.html")
	}

	err = Template(&ww, r, "non-existent.page.html", &models.TemplateData{})

	if err == nil {
		t.Error("No error in creating non-existent template")
	}
}

func TestCreateTemplateCache(t *testing.T) {
	pathToTemplate = "../../templates"

	_, err := CreateTemplateCache()

	if err != nil {
		t.Error("Error in creating template cache")
	}
}

func getRequestWithSession() (*http.Request, error) {
	r, err := http.NewRequest("GET", "/", nil)

	if err != nil {
		return nil, err
	}

	ctx := r.Context()

	// Adds session data into context
	// Load retrieves the session data for the given string token from the session store,
	// and returns a new Context containing the session data.
	// If no matching token is found, a new session will be created.
	ctx, _ = session.Load(ctx, r.Header.Get("X-Session"))

	r = r.WithContext(ctx)

	return r, nil
}