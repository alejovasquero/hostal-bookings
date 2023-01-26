package routes

import (
	"net/http"

	config "github.com/alejovasquero/hostal-bookings/internal/configs"
	"github.com/alejovasquero/hostal-bookings/internal/handler"
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
	mux.Get("/contact", handler.Repo.Contact)
	mux.Get("/rooms/jigsaw", handler.Repo.Jigsaw)
	mux.Get("/rooms/torture-premium", handler.Repo.Torture)
	mux.Get("/search-availability", handler.Repo.SearchAvalability)
	mux.Get("/make-reservation", handler.Repo.MakeReservation)

	mux.Post("/search-availability", handler.Repo.PostAvalability)

	fileServer := http.FileServer(http.Dir("./web/static"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return mux
}
