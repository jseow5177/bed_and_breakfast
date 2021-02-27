package main

import (
	"net/http"

	"github.com/justinas/nosurf"
)

/* Typical structure of a custom middleware

func MiddlewareOne(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
		fmt.Println("Middleware One")
		next.ServeHTTP(w, r) // Dispatches the request to the next handler
		fmt.Println("Middleware One Again")
	})
}

*/

// NoSurf is a middleware that prevents CSRF attacks through anti-CSRF tokens
func NoSurf(next http.Handler) http.Handler {
	// Creates a CSRF handler that calls the specified handler if the CSRF check succeeds
	csrfHandler := nosurf.New(next)

	// Sets the base cookie to use when building a CSRF token cookie
	// The CSRF token is stored in the cookie
	csrfHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true, // Cannot be accessed by client side scripts
		Path: "/", // Applies to all sits
		Secure: app.InProduction, // No HTTPS when in development
		SameSite: http.SameSiteLaxMode,
	})

	return csrfHandler
}

// SessionLoad is a middleware that automatically loads and saves session
// data for the current request, and communicates the session token to and from
// the client in a cookie.
func SessionLoad(next http.Handler) http.Handler {
	return app.Session.LoadAndSave(next)
}