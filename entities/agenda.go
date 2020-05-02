package entities

import "time"

type ProjudAgenda struct {
	ID              int64
	EscritorioID    int64
	UsuarioID       int64
	Data            time.Time
	HoraInicial     string
	HoraFinal       string
	DiaTodo         int
	Evento          string
	Obs             string
	Prazo           int
	Cloncluido      int
	Vinculo         string
	CodVinculo      string
	DataFinal       time.Time
	CodUso          string
	TipoEvento      string
	ProcessoID      int64
	ClienteID       int64
	Local           string
	SubCategoria    string
	Cor             string
	IncluirTelefone int
	UUID            string
	Excluido        int
	CodPub          int
	Privacidade     int
	FusoHorarioID   int64
	CreatedAt       time.Time
	UpdatedAt       time.Time
}
