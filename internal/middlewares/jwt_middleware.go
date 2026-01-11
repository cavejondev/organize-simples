// Package middlewares e o pacote que contem todos os middlewares do projeto
package middlewares

import (
	"net/http"
	"os"
	"strings"

	"github.com/cavejondev/organize-simples/internal/handlers"
	"github.com/cavejondev/organize-simples/internal/utils"
	"github.com/golang-jwt/jwt/v5"
)

func JWTAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		auth := r.Header.Get("Authorization")
		if auth == "" {
			utils.Error(w, http.StatusUnauthorized, handlers.CodigoUsuarioSenhaInvalida, "token não informado")
			return
		}

		parts := strings.Split(auth, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			utils.Error(w, http.StatusUnauthorized, handlers.CodigoUsuarioSenhaInvalida, "formato do token inválido")
			return
		}

		secret := os.Getenv("JWT_SECRET")
		if secret == "" {
			utils.Error(w, http.StatusInternalServerError, handlers.CodigoUsuarioSenhaInvalida, "jwt não configurado")
			return
		}

		token, err := jwt.Parse(parts[1], func(t *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})

		if err != nil || !token.Valid {
			utils.Error(w, http.StatusUnauthorized, handlers.CodigoUsuarioSenhaInvalida, "token inválido")
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			utils.Error(w, http.StatusUnauthorized, handlers.CodigoUsuarioSenhaInvalida, "claims inválidas")
			return
		}

		idUsuario, ok := claims["idUsuario"].(float64)
		if !ok {
			utils.Error(w, http.StatusUnauthorized, handlers.CodigoUsuarioSenhaInvalida, "id do usuário inválido")
			return
		}

		ctx := utils.ContextWithUsuarioID(r.Context(), int(idUsuario))
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
