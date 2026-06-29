package main

import (
	"io/fs"
	"log"
	"net/http"

	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi/v5"
	"github.com/raffy-io/bookings"
	"github.com/raffy-io/bookings/internal/handlers"
)

// Explicit parameters instead of a catch-all AppConfig struct
func routes(h *handlers.Handlers, session *scs.SessionManager, inProduction bool) http.Handler {
	mux := chi.NewRouter()

	// Inject dependencies directly into the middleware functions
	mux.Use(NoSurf(inProduction))
	mux.Use(session.LoadAndSave) // SCS SessionManager already provides this HTTP middleware directly!

	// Look how much cleaner this is. No app.Handlers. X
	mux.Get("/", h.Homepage)
	mux.Get("/about", h.About)
	mux.Get("/rooms/generals", h.Generals)
	mux.Get("/rooms/majors", h.Majors)
	mux.Get("/availability", h.Availability)
	mux.Post("/availability", h.PostAvailability)
	mux.Get("/booking", h.Booking)
	mux.Post("/booking", h.PostBooking)
	mux.Get("/booking-summary", h.BookingSummary)
	mux.Get("/contacts", h.Contacts)

	// Static assets
	staticFS, err := fs.Sub(bookings.EmbeddedAssets, "static")
	if err != nil {
		log.Fatal(err)
	}
	mux.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.FS(staticFS))))

	return mux
}