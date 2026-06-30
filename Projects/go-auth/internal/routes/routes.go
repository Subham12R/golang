package routes

import (
	"net/http"

	"go-auth/internal/handlers"
	"go-auth/internal/middleware"
)

func RegisterRoutes(
	mux *http.ServeMux,
	auth *handlers.AuthHandler,
	jwtSecret string,
) {
	mux.HandleFunc("GET /health", handlers.Health)
	mux.HandleFunc("POST /register", auth.Register)
	mux.HandleFunc("POST /login", auth.Login)
	mux.HandleFunc("POST /refresh", auth.Refresh)
	mux.HandleFunc("POST /logout", auth.Logout)


	requireAuth := middleware.Auth(jwtSecret)
	mux.Handle("GET /me", requireAuth(http.HandlerFunc(auth.Me)))
}
