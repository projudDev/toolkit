package entities

import "time"

type ProjudUsuario struct {
	ID               int64
	EscritorioID     int64
	Nome             string
	Login            string
	Senha            string
	Automatico       *int
	Tipo             *int
	UltimoAcesso     *string
	UltimoAcessoData *time.Time
	Email            *string
	Ativo            *int
	CorC             *string
	CorM             *string
	Fone             *string
	NotificaSMS      *int
	NotificaAnd      *int
	NotificaPub      *int
	Acessos          *int
	DataCadastro     *time.Time
	AgendaGoogle     *int
	Excluido         *int
	DataExclusao     *time.Time
	OAB              *string
	OABUF            *string
	OABLetra         *string
	Imagem           *string
	UpdatedAt        *time.Time
}
