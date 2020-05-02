package entities

import "time"

type ProjudNotificacao struct {
	ID           int64
	UsuarioID    int64
	EscritorioID int64
	Titulo       string
	Texto        string
	Link         string
	Lido         int
	Tipo         int
	Data         time.Time
}
