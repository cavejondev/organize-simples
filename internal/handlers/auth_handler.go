// Package handlers e o pacote que contem todos os handlers
package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/cavejondev/organize-simples/internal/services"
	"github.com/cavejondev/organize-simples/internal/utils"
)

var (
	CodigoSucesso              = "LOGIN_SUCESSO"
	CodigoUsuarioSenhaInvalida = "LOGIN_USUARIO_SENHA_INVALIDA"
	CodigoRequisicaoInvalida   = "LOGIN_REQUISICAO_INVALIDA"
	CodigoErroInterno          = "LOGIN_ERRO_INTERNO"
)

type AuthHandler struct {
	authService *services.AuthService
}

func NewAuthHandler(authService *services.AuthService) *AuthHandler {
	return &AuthHandler{authService: authService}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var req services.LoginRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.Error(w, http.StatusBadRequest, CodigoRequisicaoInvalida, "dados inválidos")
		return
	}

	resp, err := h.authService.Login(r.Context(), req)

	switch err {
	case nil:
		utils.Success(w, http.StatusOK, CodigoSucesso, resp)
	case services.ErroEmailSenhaInvalido:
		utils.Error(w, http.StatusUnauthorized, CodigoUsuarioSenhaInvalida, "email ou senha inválidos")
	default:
		utils.Error(w, http.StatusInternalServerError, CodigoErroInterno, "erro interno do servidor")
	}
}
