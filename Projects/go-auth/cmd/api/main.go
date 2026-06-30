package main

import (
	"log"
	"net/http"

	"go-auth/internal/config"
	"go-auth/internal/database"
	"go-auth/internal/handlers"
	"go-auth/internal/middleware"
	"go-auth/internal/repositories"
	"go-auth/internal/routes"
	"go-auth/internal/services"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	userRepo := repositories.NewUserRepository(db)
	refreshRepo := repositories.NewRefreshTokenRepository(db)

	authService := services.NewAuthService(
		userRepo,
		refreshRepo,
		cfg.JWTSecret,
	)

	authHandler := handlers.NewAuthHandler(authService)

	mux := http.NewServeMux()

	routes.RegisterRoutes(mux, authHandler, cfg.JWTSecret)

	log.Printf("Server running on http://localhost:%s", cfg.Port)

	if err := http.ListenAndServe(":"+cfg.Port, middleware.Logger(mux)); err != nil {
		log.Fatal(err)
	}
}
