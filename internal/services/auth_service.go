// Package services e o pacote que contem todos os serviços da aplicação
package services

import (
	"context"
	"errors"
	"os"
	"strconv"

	"github.com/cavejondev/organize-simples/internal/domain/repositories"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var (
	ErroEmailSenhaInvalido = errors.New("email ou senha inválidos")
	ErroJWTNaoConfigurado  = errors.New("JWT_SECRET não definido")
)

// LoginRequest é o modelo de request que chega ao service
type LoginRequest struct {
	Email string `json:"email"`
	Senha string `json:"senha"`
}

// LoginResponse é o modelo de resposta do service
type LoginResponse struct {
	Token string `json:"token"`
}

type AuthService struct {
	usuarioRepo repositories.UsuarioRepository
	jwtSecret   string
	jwtExpHours int
}

func NewAuthService(usuarioRepo repositories.UsuarioRepository) (*AuthService, error) {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return nil, ErroJWTNaoConfigurado
	}

	expHours := 24
	if v := os.Getenv("JWT_EXPIRE_HOURS"); v != "" {
		if h, err := strconv.Atoi(v); err == nil {
			expHours = h
		}
	}

	return &AuthService{
		usuarioRepo: usuarioRepo,
		jwtSecret:   secret,
		jwtExpHours: expHours,
	}, nil
}

// Login autentica o usuário e gera um JWT
func (s *AuthService) Login(ctx context.Context, req LoginRequest) (*LoginResponse, error) {

	usuario, err := s.usuarioRepo.BuscarPorEmail(ctx, req.Email)
	if err != nil {
		return nil, ErroEmailSenhaInvalido
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(usuario.Senha),
		[]byte(req.Senha),
	)
	if err != nil {
		return nil, ErroEmailSenhaInvalido
	}

	claims := jwt.MapClaims{
		"idUsuario": usuario.ID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	ss, err := token.SignedString([]byte(s.jwtSecret))
	if err != nil {
		return nil, err
	}

	return &LoginResponse{
		Token: ss,
	}, nil
}
