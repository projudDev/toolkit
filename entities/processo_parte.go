package entities

type ProjudProcessoParte struct {
	ID         int64 `json:"id"`
	ClienteID  int64 `json:"cliente_id"`
	ProcessoID int64 `json:"processo_id"`
	Polo       int   `json:"polo"`
}
