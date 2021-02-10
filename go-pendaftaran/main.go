package main

import (
	"fmt"
	"go-pendaftaran/router"
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"

	}
	r := router.Router()
	fmt.Println("Starting server on the port 8080 ......")
	log.Fatal(http.ListenAndServe(":"+port, r))
}
