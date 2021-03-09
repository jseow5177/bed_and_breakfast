package main

import (
	"net/http"
	"testing"
)


func TestNoSurf(t *testing.T) {
	var mh mockHandler;
	resp := NoSurf(mh);

	switch resp.(type) {
		case http.Handler:
		default:
			t.Error("Type returned by NoSurf is not http Handler")
	}
}

func TestSessionLoad(t *testing.T) {
	var mh mockHandler;
	resp := SessionLoad(mh);

	switch resp.(type) {
		case http.Handler:
		default:
			t.Error("Type returned by SessionLoad is not http Handler")
	}
}