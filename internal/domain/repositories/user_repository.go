// Package repositories e o pacote que contem todos os metodos do banco de dados, em uma interface
package repositories

import (
	"context"

	"github.com/cavejondev/organize-simples/internal/domain/models"
)

type UsuarioRepository interface {
	BuscarPorEmail(ctx context.Context, email string) (*models.Usuario, error)
}
