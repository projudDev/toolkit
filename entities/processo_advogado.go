package entities

type ProjudProcessoAdvogado struct {
	ID         int64 `json:"id"`
	AdvogadoID int64 `json:"advogado_id"`
	ProcessoID int64 `json:"processo_id"`
	Polo       int   `json:"polo"`
}
