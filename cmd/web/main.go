package main

import (
	"context"
	"encoding/gob"
	"log/slog" // 1. Add standard log/slog import
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/joho/godotenv"
	"github.com/raffy-io/bookings/internal/connection"
	"github.com/raffy-io/bookings/internal/db"
	"github.com/raffy-io/bookings/internal/handlers"
)

func main() {
	// 2. Set JSON logging right at the entrance of your app
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, nil)))

	ctx := context.Background()
    
	_ = godotenv.Load() 
	dbURL := os.Getenv("DB_URL")
	inProduction := false // Read from env or hardcode

	pool, err := connection.Connect(ctx, dbURL)
	if err != nil {
		// Change to slog.Error before killing the app
		slog.Error("Failed to initialize database", "error", err)
		os.Exit(1)
	}
	defer pool.Close() 

	queries := db.New(pool)

	// scs sessions
	gob.Register(&db.Reservation{})
	gob.Register(&db.User{})
	gob.Register(&db.Room{})
	gob.Register(&db.Restriction{})
	gob.Register(&db.RoomRestriction{})
	sessionManager := scs.New()
	sessionManager.Lifetime = 24 * time.Hour

	// Handlers get exactly what they need
	h := handlers.New(sessionManager, queries)

	// Pass ONLY the things the router actually needs to configure itself
	port := os.Getenv("PORT")
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: routes(h, sessionManager, inProduction), 
	}

	// 3. Replaced fmt.Println with a modern structured Info log!
	slog.Info("Server is running", "addr", srv.Addr)

	err = srv.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		slog.Error("Server failed to start", "error", err)
		os.Exit(1)
	}
}