package entities

type ProjudProcessoCliente struct {
	ID            int64     `json:"id"`
	EscritorioID  int64     `json:"escritorio_id"`
	Nome          string    `json:"nome"`
	DataCad       time.Time `json:"data_cad"`
	Classificacao int       `json:"classificacao"`
}
