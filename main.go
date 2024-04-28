package main

import (
	"database/sql"
	"log"
	"michafdlr/blog_aggregator/internal/database"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")
	Conn := os.Getenv("CONN")

	db, err := sql.Open("postgres", Conn)

	if err != nil {
		log.Fatal(err)
	}

	dbQueries := database.New(db)

	apiCfg := apiConfig{
		DB: dbQueries,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /v1/readiness", ReadinessHandler)
	mux.HandleFunc("GET /v1/error", ErrorHandler)
	mux.HandleFunc("POST /v1/users", apiCfg.CreateUserHandler)
	mux.HandleFunc("GET /v1/users", apiCfg.middlewareAuth(apiCfg.GetCurrentUserHandler))
	mux.HandleFunc("POST /v1/feeds", apiCfg.middlewareAuth(apiCfg.CreateFeedHandler))

	corsMux := corsMiddleware(mux)
	srv := &http.Server{
		Handler: corsMux,
		Addr:    ":" + port,
	}
	log.Printf("Serving on port %s", port)
	log.Fatal(srv.ListenAndServe())
}

type apiConfig struct {
	DB *database.Queries
}
