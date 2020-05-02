package entities

import "time"

type ProjudProcessoAndamento struct {
	ID          int64     `json:"id"`
	ProcessoID  int64     `json:"processo_id"`
	Descricao   string    `json:"descricao"`
	Data        time.Time `json:"data"`
	DataCad     time.Time `json:"data_cad"`
	AndamentoID int64     `json:"andamento_id"`
	Enviado     int       `json:"enviado"`
}
