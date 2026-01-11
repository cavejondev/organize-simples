package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/cavejondev/organize-simples/internal/db"
	"github.com/cavejondev/organize-simples/internal/handlers"
	infraRepo "github.com/cavejondev/organize-simples/internal/repositories"
	"github.com/cavejondev/organize-simples/internal/services"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

func main() {
	// Carrega variáveis de ambiente
	if err := godotenv.Load(); err != nil {
		log.Println("arquivo .env não encontrado, usando variáveis do sistema")
	}

	// Porta
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Conexão com Postgres
	ctx := context.Background()
	dbPool, err := db.NewPostgresPool(ctx)
	if err != nil {
		log.Fatal("erro ao conectar no banco:", err)
	}
	defer dbPool.Close()

	// Repositories (infra)
	usuarioRepo := infraRepo.NewUsuarioRepositoryPg(dbPool)

	// Services
	authService, err := services.NewAuthService(usuarioRepo)
	if err != nil {
		log.Fatal("erro ao criar AuthService:", err)
	}

	// Handlers
	authHandler := handlers.NewAuthHandler(authService)

	// Router
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Rotas
	r.Post("/login", authHandler.Login)

	log.Println("Servidor rodando na porta", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
