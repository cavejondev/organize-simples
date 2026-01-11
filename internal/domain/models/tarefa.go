package models

import "time"

type Tarefa struct {
	ID            int
	IDUsuario     int
	Titulo        string
	Descricao     string
	Status        string
	DataAgendada  *time.Time
	DataConclusao *time.Time
	CreatedAt     time.Time
}
