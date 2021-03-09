package main

import (
	"testing"

	"github.com/go-chi/chi"
)


func TestRoutes(t *testing.T) {
	mux := routes()

	switch mux.(type) {
		case *chi.Mux:
		default:
			t.Error("Type returned by routes() is not *chi.Mux")
	}
}