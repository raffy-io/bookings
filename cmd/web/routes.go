package main

import (
	"io/fs"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/raffy-io/bookings"
)

func (app *AppConfig) routes() http.Handler {
	mux := chi.NewRouter()
	mux.Use(middleware.Logger)
	mux.Use(app.NoSurf)
	mux.Use(app.SessionLoad)

	mux.Get("/", app.Handlers.Home)
	mux.Get("/about", app.Handlers.About)
	mux.Get("/generals-quarters", app.Handlers.Generals)
	mux.Get("/majors-suite", app.Handlers.Majors)
	mux.Get("/search-availability",app.Handlers.Availability)
	mux.Post("/search-availability",app.Handlers.PostAvailability)
	mux.Get("/make-reservation",app.Handlers.Reservation)
	mux.Post("/make-reservation",app.Handlers.PostReservation)
	mux.Get("/contacts",app.Handlers.Contacts)

	// static assets
	staticFS, err := fs.Sub(bookings.EmbeddedAssets, "static")
	if err != nil {
		log.Fatal(err)
	}
	mux.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.FS(staticFS))))

	return mux
}
