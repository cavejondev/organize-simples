// Package services e o pacote que contem todos os serviços da aplicação
package services

import (
	"errors"

	"github.com/cavejondev/organize-simples/internal/repositories"
	"golang.org/x/crypto/bcrypt"
)

// Tipos de request e response
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Message string `json:"message"`
}

// Erros
var (
	ErroEmailSenhaInvalido = errors.New("email ou senha inválidos")
)

// Login autentica o usuário
func Login(req LoginRequest) (*LoginResponse, error) {
	user, err := repositories.FindUserByEmail(req.Email)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, ErroEmailSenhaInvalido
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Senha), []byte(req.Password))
	if err != nil {
		return nil, ErroEmailSenhaInvalido
	}

	return &LoginResponse{Message: "login ok"}, nil
}
