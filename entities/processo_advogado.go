package entities

type ProjudProcessoAdvogado struct {
	ID         int64  `json:"id"`
	AdvogadoID int64  `json:"advogado_id"`
	ProcessoID int64  `json:"processo_id"`
	Nome       string `json:"nome"`
	Polo       int    `json:"polo"`
}
