package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/cavejondev/organize-simples/internal/db"
	"github.com/cavejondev/organize-simples/internal/handlers"
	"github.com/cavejondev/organize-simples/internal/middlewares"
	infraRepo "github.com/cavejondev/organize-simples/internal/repositories"
	"github.com/cavejondev/organize-simples/internal/services"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

func main() {
	// ENV
	if err := godotenv.Load(); err != nil {
		log.Println("arquivo .env não encontrado, usando variáveis do sistema")
	}

	// Porta
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// DB
	ctx := context.Background()
	dbPool, err := db.NewPostgresPool(ctx)
	if err != nil {
		log.Fatal("erro ao conectar no banco:", err)
	}
	defer dbPool.Close()

	usuarioRepo := infraRepo.NewUsuarioRepositoryPg(dbPool)
	tarefaRepo := infraRepo.NewTarefaRepositoryPg(dbPool)

	authService, err := services.NewAuthService(usuarioRepo)
	if err != nil {
		log.Fatal("erro ao criar AuthService:", err)
	}

	tarefaService := services.NewTarefaService(tarefaRepo)

	authHandler := handlers.NewAuthHandler(authService)
	tarefaHandler := handlers.NewTarefaHandler(tarefaService)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Pública
	r.Post("/login", authHandler.Login)

	// Protegidas
	r.Route("/tarefa", func(rt chi.Router) {
		rt.Use(middlewares.JWTAuth)

		rt.Post("/", tarefaHandler.Criar)
		rt.Get("/", tarefaHandler.Listar)
		rt.Put("/{id}", tarefaHandler.Atualizar)
	})

	log.Println("Servidor rodando na porta", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
