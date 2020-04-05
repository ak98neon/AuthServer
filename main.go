package main

import (
	"github.com/ak98neon/authserver/route"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	e := godotenv.Load()

	if e != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")

	// Handle routes
	http.Handle("/", route.Routes())

	// serve
	log.Printf("Server up on port '%s'", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
