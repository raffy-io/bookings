package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/raffy-io/bookings/internal/handlers"
)

type AppConfig struct {
	InProduction bool
	Session      *scs.SessionManager
	Handlers     *handlers.Handlers
}

func main() {
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
