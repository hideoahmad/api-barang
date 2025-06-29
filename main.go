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

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080" // Default port lokal
    }

    fmt.Println("API running on port", port)
    log.Fatal(http.ListenAndServe(":"+port, router))
}
