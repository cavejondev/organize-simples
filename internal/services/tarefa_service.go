// Package services contem todos os serviços da aplicação
package services

import (
	"context"
	"errors"
	"time"

	"github.com/cavejondev/organize-simples/internal/domain/models"
	"github.com/cavejondev/organize-simples/internal/domain/repositories"
)

var (
	ErroStatusInvalido = errors.New("status inválido")
)

// TarefaRequest é o modelo que o handler envia para o service
type TarefaRequest struct {
	Titulo        string     `json:"titulo"`
	Descricao     string     `json:"descricao"`
	Status        string     `json:"status"`        // 'A', 'F', 'C'
	DataAgendada  *time.Time `json:"dataAgendada"`  // opcional
	DataConclusao *time.Time `json:"dataConclusao"` // opcional
}

// TarefaResponse é o modelo que o service devolve para o handler
type TarefaResponse struct {
	ID            int        `json:"id"`
	Titulo        string     `json:"titulo"`
	Descricao     string     `json:"descricao"`
	Status        string     `json:"status"`
	DataAgendada  *time.Time `json:"dataAgendada,omitempty"`
	DataConclusao *time.Time `json:"dataConclusao,omitempty"`
	CreatedAt     time.Time  `json:"createdAt"`
}

type TarefaService struct {
	repo repositories.TarefaRepository
}

func NewTarefaService(repo repositories.TarefaRepository) *TarefaService {
	return &TarefaService{repo: repo}
}

// Criar cria uma nova tarefa
func (s *TarefaService) Criar(ctx context.Context, idUsuario int, req *TarefaRequest) (*TarefaResponse, error) {
	t := &models.Tarefa{
		IDUsuario:    idUsuario,
		Titulo:       req.Titulo,
		Descricao:    req.Descricao,
		Status:       "A",
		DataAgendada: req.DataAgendada,
	}
	if err := s.repo.Criar(ctx, t); err != nil {
		return nil, err
	}
	return &TarefaResponse{
		ID:            t.ID,
		Titulo:        t.Titulo,
		Descricao:     t.Descricao,
		Status:        t.Status,
		DataAgendada:  t.DataAgendada,
		DataConclusao: t.DataConclusao,
		CreatedAt:     t.CreatedAt,
	}, nil
}

// Listar retorna todas as tarefas do usuário
func (s *TarefaService) Listar(ctx context.Context, idUsuario int) ([]TarefaResponse, error) {
	tarefas, err := s.repo.ListarPorUsuario(ctx, idUsuario)
	if err != nil {
		return nil, err
	}
	var resp []TarefaResponse
	for _, t := range tarefas {
		resp = append(resp, TarefaResponse{
			ID:            t.ID,
			Titulo:        t.Titulo,
			Descricao:     t.Descricao,
			Status:        t.Status,
			DataAgendada:  t.DataAgendada,
			DataConclusao: t.DataConclusao,
			CreatedAt:     t.CreatedAt,
		})
	}
	return resp, nil
}

// Atualizar altera uma tarefa existente
func (s *TarefaService) Atualizar(ctx context.Context, idUsuario int, idTarefa int, req *TarefaRequest) (*TarefaResponse, error) {
	if req.Status != "A" && req.Status != "F" && req.Status != "C" {
		return nil, ErroStatusInvalido
	}
	t := &models.Tarefa{
		ID:            idTarefa,
		IDUsuario:     idUsuario,
		Titulo:        req.Titulo,
		Descricao:     req.Descricao,
		Status:        req.Status,
		DataAgendada:  req.DataAgendada,
		DataConclusao: req.DataConclusao,
	}
	if err := s.repo.Atualizar(ctx, t); err != nil {
		return nil, err
	}
	return &TarefaResponse{
		ID:            t.ID,
		Titulo:        t.Titulo,
		Descricao:     t.Descricao,
		Status:        t.Status,
		DataAgendada:  t.DataAgendada,
		DataConclusao: t.DataConclusao,
		CreatedAt:     t.CreatedAt,
	}, nil
}

// Deletar remove uma tarefa
func (s *TarefaService) Deletar(ctx context.Context, idUsuario int, idTarefa int) error {
	return s.repo.Deletar(ctx, idTarefa, idUsuario)
}
