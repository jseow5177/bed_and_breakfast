package render

import (
	"encoding/gob"
	"net/http"
	"os"
	"testing"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/jseow5177/bed_and_breakfast/internals/config"
	"github.com/jseow5177/bed_and_breakfast/internals/models"
)

var session *scs.SessionManager
var testApp config.AppConfig

func TestMain(m *testing.M) {
	gob.Register(models.Reservation{})

	testApp.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = false

	testApp.Session = session

	app = &testApp

	os.Exit(m.Run())
}

// Create a mock writer
type mockWriter struct{}
func (ww *mockWriter) WriteHeader(statusCode int) {}
func (ww *mockWriter) Write(b []byte) (int, error) { return len(b), nil }
func (ww *mockWriter) Header() http.Header { return http.Header{} }