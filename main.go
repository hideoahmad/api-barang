package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"my-api/config"
	"my-api/routes"
)

func main() {
    config.Connect()
    router := routes.SetupRoutes()

    // Middleware CORS
    corsHandler := func(h http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            w.Header().Set("Access-Control-Allow-Origin", "*")
            w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
            w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

            // Handle preflight
            if r.Method == "OPTIONS" {
                w.WriteHeader(http.StatusOK)
                return
            }

            h.ServeHTTP(w, r)
        })
    }

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
    }

    fmt.Println("API running on port", port)
    log.Fatal(http.ListenAndServe(":"+port, corsHandler(router)))
}
