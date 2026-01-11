// Package repositories e o pacote que contem todos os comando de banco de dados
package repositories

import (
	"context"

	"github.com/cavejondev/organize-simples/internal/domain/models"
	domainRepo "github.com/cavejondev/organize-simples/internal/domain/repositories"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UsuarioRepositoryPg struct {
	db *pgxpool.Pool
}

func NewUsuarioRepositoryPg(db *pgxpool.Pool) domainRepo.UsuarioRepository {
	return &UsuarioRepositoryPg{db: db}
}

func (r *UsuarioRepositoryPg) BuscarPorEmail(ctx context.Context, email string) (*models.Usuario, error) {
	row := r.db.QueryRow(
		ctx,
		`SELECT id, email, senha FROM usuarios WHERE email = $1`,
		email,
	)

	var u models.Usuario
	if err := row.Scan(&u.ID, &u.Email, &u.Senha); err != nil {
		return nil, err
	}

	return &u, nil
}
