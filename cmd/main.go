package main

import (
	"log"
	"net/http"
	"os"

	"github.com/cavejondev/organize-simples/internal/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/joho/godotenv"
)

func main() {
	// Carrega .env
	if err := godotenv.Load(); err != nil {
		log.Println("arquivo .env não encontrado, usando variáveis do sistema")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	r := chi.NewRouter()

	// Middlewares básicos
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Rotas
	r.Post("/login", handlers.Login)

	log.Println("Servidor rodando na porta", port)
	log.Fatal(http.ListenAndServe(":"+port, r))
}
