package utils

import (
	"context"
	"errors"
)

var ErroUsuarioNaoAutenticado = errors.New("usuario não autenticado")

type contextKey string

const usuarioIDKey contextKey = "idusuario"

func ContextWithUsuarioID(ctx context.Context, idUsuario int) context.Context {
	return context.WithValue(ctx, usuarioIDKey, idUsuario)
}

func UsuarioIDFromContext(ctx context.Context) (int, error) {
	id, ok := ctx.Value(usuarioIDKey).(int)
	if !ok {
		return 0, ErroUsuarioNaoAutenticado
	}
	return id, nil
}
