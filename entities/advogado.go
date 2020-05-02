package entities

type ProjudAdvogado struct {
	ID           int64  `json:"id"`
	EscritorioID int64  `json:"escritorio_id"`
	Nome         string `json:"nome"`
}
