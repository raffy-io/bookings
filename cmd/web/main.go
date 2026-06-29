package main

import (
	"context"
	"encoding/gob"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/joho/godotenv"
	"github.com/raffy-io/bookings/internal/connection"
	"github.com/raffy-io/bookings/internal/db"
	"github.com/raffy-io/bookings/internal/handlers"
	"github.com/raffy-io/bookings/internal/models"
)




func main() {
    ctx := context.Background()
    
    _ = godotenv.Load() 
    dbURL := os.Getenv("DB_URL")
    inProduction := false // Read from env or hardcode

    pool, err := connection.Connect(ctx, dbURL)
    if err != nil {
        log.Fatalf("Failed to initialize database: %v\n", err)
    }
    defer pool.Close() 

    queries := db.New(pool)

    gob.Register(models.ReservationSummary{})
    sessionManager := scs.New()
    sessionManager.Lifetime = 24 * time.Hour

    // Handlers get exactly what they need
    h := handlers.New(sessionManager, queries)

    // Pass ONLY the things the router actually needs to configure itself
    srv := &http.Server{
        Addr:    ":8080",
        Handler: routes(h, sessionManager, inProduction), 
    }

   fmt.Println("Server is running on port", srv.Addr)
	err = srv.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		log.Fatalf("Server failed to start: %v", err)
	}
}