// Package handlers e o pacote que contem todos os handlers
package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/cavejondev/organize-simples/internal/repositories"
	"github.com/cavejondev/organize-simples/internal/services"
	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "dados inválidos", http.StatusBadRequest)
		return
	}

	user, err := repositories.FindUserByEmail(req.Email)
	if err != nil {
		http.Error(w, services.ErroEmailSenhaInvalido.Error(), http.StatusUnauthorized)
		return
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Senha),
		[]byte(req.Password),
	)
	if err != nil {
		http.Error(w, services.ErroEmailSenhaInvalido.Error(), http.StatusUnauthorized)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"login ok"}`))
}
