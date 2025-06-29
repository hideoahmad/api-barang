package main

import (
	"fmt"
	"log"
	"net/http"

	"my-api/config"
	"my-api/routes"
)

func main() {
    config.Connect()
    router := routes.SetupRoutes()
    fmt.Println("API running on port 8080")
    log.Fatal(http.ListenAndServe(":8080", router))
}
