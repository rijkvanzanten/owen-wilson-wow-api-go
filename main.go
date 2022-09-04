package main

import (
	"errors"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {
	envErr := godotenv.Load()

	if envErr != nil {
		log.Fatalf("Can't read environment variables from disk\n")
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", RouteHome)
	mux.HandleFunc("/movies", RouteMovies)
	handler := cors.Default().Handler(mux)
	port := ":" + os.Getenv("PORT")
	listenAndServeErr := http.ListenAndServe(port, handler)

	if errors.Is(listenAndServeErr, http.ErrServerClosed) {
		log.Printf("Server closed\n")
	} else {
		log.Fatalf("Error starting server: %s\n", listenAndServeErr)
	}
}