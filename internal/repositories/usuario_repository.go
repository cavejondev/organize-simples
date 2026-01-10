// Package repository e o pacote que contem todos os comando de banco de dados
package repository

import (
	"context"

	"github.com/cavejondev/organize-simples/internal/db"
	"github.com/cavejondev/organize-simples/internal/models"
)

func FindUserByEmail(email string) (*models.Usuario, error) {
	row := db.Pool.QueryRow(
		context.Background(),
		`SELECT id, email, senha FROM users WHERE email = $1`,
		email,
	)

	var u models.Usuario
	err := row.Scan(&u.ID, &u.Email, &u.Senha)
	if err != nil {
		return nil, err
	}

	return &u, nil
}
