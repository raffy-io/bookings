package main

import (
	"encoding/gob"
	"fmt"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/raffy-io/bookings/internal/handlers"
	"github.com/raffy-io/bookings/internal/models"
)

type AppConfig struct {
	InProduction bool
	Session      *scs.SessionManager
	Handlers     *handlers.Handlers
}

func main() {
	gob.Register(models.ReservationSummary{})
	sessionManager := scs.New()
	sessionManager.Lifetime = 24 * time.Hour

	h := handlers.New(sessionManager)

	app := &AppConfig{
		InProduction: false,
		Session:      sessionManager,
		Handlers:     h,
	}

	srv := &http.Server{
		Addr:    ":8080",
		Handler: app.routes(),
	}

	fmt.Println("Server is running on port", srv.Addr)
	srv.ListenAndServe()
}
