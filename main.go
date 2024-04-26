package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")

	mux := http.NewServeMux()
	mux.HandleFunc("GET /v1/readiness", ReadinessHandler)
	mux.HandleFunc("GET /v1/error", ErrorHandler)

	corsMux := middlewareCors(mux)
	srv := &http.Server{
		Handler: corsMux,
		Addr:    ":" + port,
	}
	log.Printf("Serving on port %s", port)
	log.Fatal(srv.ListenAndServe())
}
