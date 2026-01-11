package repositories

import (
	"context"

	"github.com/cavejondev/organize-simples/internal/domain/models"
)

type TarefaRepository interface {
	Criar(ctx context.Context, t *models.Tarefa) error
	ListarPorUsuario(ctx context.Context, idUsuario int) ([]models.Tarefa, error)
	BuscarPorID(ctx context.Context, id int, idUsuario int) (*models.Tarefa, error)
	Atualizar(ctx context.Context, t *models.Tarefa) error
	Deletar(ctx context.Context, id int, idUsuario int) error
}
