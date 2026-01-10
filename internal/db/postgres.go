// Package db e o pacote que contem o banco de dados
package db

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var Pool *pgxpool.Pool

func Connect() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		return fmt.Errorf("DATABASE_URL não definida")
	}

	pool, err := pgxpool.New(ctx, dbURL)
	if err != nil {
		return err
	}

	// testa conexão
	if err := pool.Ping(ctx); err != nil {
		return err
	}

	Pool = pool
	return nil
}
