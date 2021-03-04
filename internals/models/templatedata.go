package models

import "github.com/jseow5177/bed_and_breakfast/internals/forms"

// TemplateData holds data sent from handlers to templates
type TemplateData struct {
	StringMap map[string]string
	IntMap map[string]int
	FloatMap map[string]float64
	Data map[string]interface{} // Data of any type
	CSRFToken string
	Flash string
	Warning string
	Error string
	Form *forms.Form
}