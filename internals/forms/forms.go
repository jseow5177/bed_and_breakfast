package forms

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/asaskevich/govalidator"
)

// Form creates a custom form struct, embeds a url.Values object and custom errors struct
type Form struct {
	Values url.Values
	Errors errors
}

// New initialzes a form struct
func New(data url.Values) *Form {
	return &Form {
		Values: data,
		Errors: make(errors), // Empty errors map
	}
}

// Required checks if a list of required fields are empty
func (f *Form)Required(fields ...string) {
	for _, field := range fields {
		value := f.Values.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "This field cannot be blank")
		}
	}
}

// MinLength checks for string minimum length
func (f *Form)MinLength(field string, length int) {
	value := f.Values.Get(field)
	if len(value) < length {
		f.Errors.Add(field, fmt.Sprintf("This field must be at least %d characters long", length))
	}
}

// IsEmail checks if field has a valid email value
func (f *Form)IsEmail(field string) {
	value := f.Values.Get(field)
	if !govalidator.IsEmail(value) {
		f.Errors.Add(field, "Email is invalid")
	}
}

// Has returns false if form value of a field is empty, else returns true
func (f *Form)Has(field string, r *http.Request) bool {
	x := r.Form.Get(field)
	if x == "" {
		return false
	}
	return true
}

// Valid returns true if there are no errors in the form, else false
func (f *Form)Valid() bool {
	return len(f.Errors) == 0
}