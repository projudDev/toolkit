package entities

type ProjudComarca struct {
	ID       int64  `json:"id"`
	Orgao    int    `json:"orgao"`
	Tribunal int    `json:"tribunal"`
	Numero   string `json:"numero"`
	Comarca  string `json:"comarca"`
	Local    string `json:"local"`
	UF       string `json:"uf"`
}
