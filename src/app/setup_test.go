package main

import (
	"log"
	"net/http"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
		log.Println("Test setup...")
    exitVal := m.Run()
    log.Println("Test teardown...")
    os.Exit(exitVal)
}

type mockHandler struct {}
func (h mockHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {}