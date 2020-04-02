package entities

type ProjudProcessoParte struct {
	ID                 int64  `json:"id"`
	ClienteID          int64  `json:"cliente_id"`
	Nome               string `json:"nome"`
	ProcessoID         int64  `json:"processo_id"`
	Polo               int    `json:"polo"`
	IsClientePrincipal bool   `json:"is_cliente_principal"`
}
