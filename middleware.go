package main

import (
	"michafdlr/blog_aggregator/internal/database"
	"net/http"
	"strings"
)

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if origin != "" {
			if strings.HasPrefix(origin, "http://") || strings.HasPrefix(origin, "https://") {
				w.Header().Set("Access-Control-Allow-Origin", origin)
			}

			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "*")
			w.Header().Set("Access-Control-Expose-Headers", "Link")
			w.Header().Set("Access-Control-Max-Age", "300")

			if r.Method == http.MethodOptions {
				w.WriteHeader(http.StatusOK)
				return
			}
		}

		next.ServeHTTP(w, r)
	})
}

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (cfg *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := GetHeader(r.Header)
		if err != nil {
			respondWithError(w, http.StatusUnauthorized, "apikey missing")
			return
		}
		user, err := cfg.DB.GetUserByKey(r.Context(), apiKey)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Couldn't find user")
			return
		}
		handler(w, r, user)
	}
}
