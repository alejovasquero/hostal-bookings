package routes

import (
	"net/http"

	"github.com/alejovasquero/hostal-bookings/config"
	"github.com/alejovasquero/hostal-bookings/pkg/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func RoutesWithChi(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(middleware.Logger)
	mux.Use(WriteToConsole)
	mux.Use(NoSurf(*app))
	mux.Use(SessionLoad(*app))

	mux.Get("/", handler.Repo.Index)
	mux.Get("/about", handler.Repo.About)

	fileServer := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
