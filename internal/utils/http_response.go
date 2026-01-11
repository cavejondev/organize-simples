// Package utils e o pacote que contem todos os metodos utils do projeto
package utils

import (
	"encoding/json"
	"net/http"
)

// Response padrão para JSON
type Response struct {
	Codigo   string      `json:"codigo"`
	Mensagem string      `json:"mensagem"`
	Dados    interface{} `json:"dados"`
}

// JSON envia uma resposta HTTP com status e payload
func JSON(w http.ResponseWriter, status int, payload Response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}

// Error envia uma resposta de erro HTTP padronizada
func Error(
	w http.ResponseWriter,
	status int,
	codigo string,
	mensagem string,
) {
	JSON(w, status, Response{
		Codigo:   codigo,
		Mensagem: mensagem,
		Dados:    nil,
	})
}

// Success envia uma resposta de sucesso HTTP padronizada
func Success(
	w http.ResponseWriter,
	status int,
	codigo string,
	dados interface{},
) {
	JSON(w, status, Response{
		Codigo:   codigo,
		Mensagem: "",
		Dados:    dados,
	})
}
