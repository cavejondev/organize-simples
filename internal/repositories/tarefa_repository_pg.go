// Package repositories e o pacote que contem todos os comandos de banco de dados
package repositories

import (
	"context"

	"github.com/cavejondev/organize-simples/internal/domain/models"
	domainRepo "github.com/cavejondev/organize-simples/internal/domain/repositories"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TarefaRepositoryPg struct {
	db *pgxpool.Pool
}

func NewTarefaRepositoryPg(db *pgxpool.Pool) domainRepo.TarefaRepository {
	return &TarefaRepositoryPg{db: db}
}

func (r *TarefaRepositoryPg) Criar(ctx context.Context, t *models.Tarefa) error {
	return r.db.QueryRow(ctx, `
		INSERT INTO tarefa (
			idusuario,
			titulo,
			descricao,
			status,
			dataagendada
		)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, createdat
	`,
		t.IDUsuario,
		t.Titulo,
		t.Descricao,
		t.Status,
		t.DataAgendada,
	).Scan(&t.ID, &t.CreatedAt)
}

func (r *TarefaRepositoryPg) ListarPorUsuario(ctx context.Context, idUsuario int) ([]models.Tarefa, error) {
	rows, err := r.db.Query(ctx, `
		SELECT
			id,
			titulo,
			descricao,
			status,
			dataagendada,
			dataconclusao,
			createdat
		FROM tarefa
		WHERE idusuario = $1
		ORDER BY id
	`, idUsuario)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tarefas []models.Tarefa

	for rows.Next() {
		var t models.Tarefa
		if err := rows.Scan(
			&t.ID,
			&t.Titulo,
			&t.Descricao,
			&t.Status,
			&t.DataAgendada,
			&t.DataConclusao,
			&t.CreatedAt,
		); err != nil {
			return nil, err
		}

		t.IDUsuario = idUsuario
		tarefas = append(tarefas, t)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tarefas, nil
}

func (r *TarefaRepositoryPg) BuscarPorID(ctx context.Context, id int, idUsuario int) (*models.Tarefa, error) {
	var t models.Tarefa

	err := r.db.QueryRow(ctx, `
		SELECT
			id,
			titulo,
			descricao,
			status,
			dataagendada,
			dataconclusao,
			createdat
		FROM tarefa
		WHERE id = $1 AND idusuario = $2
	`, id, idUsuario).Scan(
		&t.ID,
		&t.Titulo,
		&t.Descricao,
		&t.Status,
		&t.DataAgendada,
		&t.DataConclusao,
		&t.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	t.IDUsuario = idUsuario
	return &t, nil
}

func (r *TarefaRepositoryPg) Atualizar(ctx context.Context, t *models.Tarefa) error {
	_, err := r.db.Exec(ctx, `
		UPDATE tarefa
		SET
			titulo = $1,
			descricao = $2,
			status = $3,
			dataagendada = $4,
			dataconclusao = $5
		WHERE id = $6 AND idusuario = $7
	`,
		t.Titulo,
		t.Descricao,
		t.Status,
		t.DataAgendada,
		t.DataConclusao,
		t.ID,
		t.IDUsuario,
	)

	return err
}

func (r *TarefaRepositoryPg) Deletar(ctx context.Context, id int, idUsuario int) error {
	_, err := r.db.Exec(ctx, `
		DELETE FROM tarefa
		WHERE id = $1 AND idusuario = $2
	`, id, idUsuario)

	return err
}
