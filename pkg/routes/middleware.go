package routes

import (
	"fmt"
	"net/http"

	config "github.com/alejovasquero/hostal-bookings/internal/configs"
	"github.com/justinas/nosurf"
)

func WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Hit the page")
		next.ServeHTTP(w, r)
	})
}

func NoSurf(config config.AppConfig) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		csrfHandler := nosurf.New(next)
		csrfHandler.SetFailureHandler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("CSRF FAILURE")
			w.WriteHeader(http.StatusForbidden)
		}))

		csrfHandler.SetBaseCookie(http.Cookie{
			HttpOnly: false,
			Path:     "/",
			Secure:   config.InProduction,
			SameSite: http.SameSiteLaxMode,
		})

		return csrfHandler
	}
}

func SessionLoad(config config.AppConfig) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return config.Session.LoadAndSave(next)
	}
}
