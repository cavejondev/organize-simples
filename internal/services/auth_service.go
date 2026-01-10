// Package services e o pacote que contem todos os serviços da aplicação
package services

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/cavejondev/organize-simples/internal/repositories"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// LoginRequest é o modelo de request deve chegar ao service
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// LoginResponse é o modelo de resposta que deve ser retornado pelo service
type LoginResponse struct {
	Token string `json:"token"`
}

// Erros
var (
	ErroEmailSenhaInvalido = errors.New("email ou senha inválidos")
	ErroJWTNaoConfigurado  = errors.New("JWT_SECRET não definido")
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

	// Criar token JWT
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return nil, ErroJWTNaoConfigurado
	}

	expHours := 24 // padrão
	if v := os.Getenv("JWT_EXPIRE_HOURS"); v != "" {
		// converte string para int
		fmt.Sscanf(v, "%d", &expHours)
	}

	claims := jwt.MapClaims{
		"idUsuario": user.ID,
		"exp":       time.Now().Add(time.Hour * time.Duration(expHours)).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(secret))
	if err != nil {
		return nil, err
	}

	return &LoginResponse{Token: ss}, nil
}
