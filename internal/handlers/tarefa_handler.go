// Package handlers contem todos os handlers da aplicação
package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/cavejondev/organize-simples/internal/services"
	"github.com/cavejondev/organize-simples/internal/utils"
	"github.com/go-chi/chi/v5"
)

var (
	CodigoTarefaCriada         = "TAREFA_CRIADA"
	CodigoTarefaListada        = "TAREFA_LISTADA"
	CodigoTarefaAtualizada     = "TAREFA_ATUALIZADA"
	CodigoTarefaDeletada       = "TAREFA_DELETADA"
	CodigoTarefaRequisicaoErro = "TAREFA_REQUISICAO_INVALIDA"
	CodigoTarefaNaoEncontrada  = "TAREFA_NAO_ENCONTRADA"
	CodigoTarefaErroInterno    = "TAREFA_ERRO_INTERNO"
	CodigoNaoAutenticado       = "NAO_AUTENTICADO"
)

type TarefaHandler struct {
	service *services.TarefaService
}

func NewTarefaHandler(service *services.TarefaService) *TarefaHandler {
	return &TarefaHandler{service: service}
}

// Criar cria uma nova tarefa
func (h *TarefaHandler) Criar(w http.ResponseWriter, r *http.Request) {
	idUsuario, err := utils.UsuarioIDFromContext(r.Context())
	if err != nil {
		utils.Error(w, http.StatusUnauthorized, CodigoNaoAutenticado, err.Error())
		return
	}

	var req services.TarefaRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.Error(w, http.StatusBadRequest, CodigoTarefaRequisicaoErro, "dados inválidos")
		return
	}

	resp, err := h.service.Criar(r.Context(), idUsuario, &req)
	if err != nil {
		utils.Error(w, http.StatusInternalServerError, CodigoTarefaErroInterno, "erro ao criar tarefa: "+err.Error())
		return
	}

	utils.Success(w, http.StatusCreated, CodigoTarefaCriada, resp)
}

// Listar retorna todas as tarefas do usuário
func (h *TarefaHandler) Listar(w http.ResponseWriter, r *http.Request) {
	idUsuario, err := utils.UsuarioIDFromContext(r.Context())
	if err != nil {
		utils.Error(w, http.StatusUnauthorized, CodigoNaoAutenticado, err.Error())
		return
	}

	resp, err := h.service.Listar(r.Context(), idUsuario)
	if err != nil {
		utils.Error(w, http.StatusInternalServerError, CodigoTarefaErroInterno, "erro ao listar tarefas: "+err.Error())
		return
	}

	utils.Success(w, http.StatusOK, CodigoTarefaListada, resp)
}

// Atualizar altera uma tarefa existente
func (h *TarefaHandler) Atualizar(w http.ResponseWriter, r *http.Request) {
	idUsuario, err := utils.UsuarioIDFromContext(r.Context())
	if err != nil {
		utils.Error(w, http.StatusUnauthorized, CodigoNaoAutenticado, err.Error())
		return
	}

	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		utils.Error(w, http.StatusBadRequest, CodigoTarefaRequisicaoErro, "id inválido")
		return
	}

	var req services.TarefaRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.Error(w, http.StatusBadRequest, CodigoTarefaRequisicaoErro, "dados inválidos")
		return
	}

	resp, err := h.service.Atualizar(r.Context(), idUsuario, id, &req)
	if err != nil {
		if err == services.ErroStatusInvalido {
			utils.Error(w, http.StatusBadRequest, CodigoTarefaRequisicaoErro, err.Error())
			return
		}
		utils.Error(w, http.StatusInternalServerError, CodigoTarefaErroInterno, err.Error())
		return
	}

	utils.Success(w, http.StatusOK, CodigoTarefaAtualizada, resp)
}

// Deletar remove uma tarefa
func (h *TarefaHandler) Deletar(w http.ResponseWriter, r *http.Request) {
	idUsuario, err := utils.UsuarioIDFromContext(r.Context())
	if err != nil {
		utils.Error(w, http.StatusUnauthorized, CodigoNaoAutenticado, err.Error())
		return
	}

	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		utils.Error(w, http.StatusBadRequest, CodigoTarefaRequisicaoErro, "id inválido")
		return
	}

	if err := h.service.Deletar(r.Context(), idUsuario, id); err != nil {
		utils.Error(w, http.StatusInternalServerError, CodigoTarefaErroInterno, err.Error())
		return
	}

	utils.Success(w, http.StatusOK, CodigoTarefaDeletada, nil)
}
