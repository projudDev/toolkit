package entities

import "time"

type ProjudProcesso struct {
	ID                 int64                     `json:"id"`
	EscritorioID       int64                     `json:"escritorio_id"`
	ClienteID          int64                     `json:"cliente_id"`
	Autos              string                    `json:"autos"`
	Nome               *string                   `json:"nome"`
	Comarca            *string                   `json:"comarca"`
	Vara               *string                   `json:"vara"`
	Acao               *string                   `json:"acao"`
	Protocolo          *string                   `json:"protocolo"`
	ValorCausa         *string                   `json:"valor_causa"`
	Honorarios         *string                   `json:"honorarios"`
	ParteContra        *string                   `json:"parte_contra"`
	Situacao           *string                   `json:"situacao"`
	Obs                *string                   `json:"obs"`
	Ativo              *bool                     `json:"ativo"`
	MotivoEncerramento *bool                     `json:"motivo_encerramento"`
	Recurso            *bool                     `json:"recurso"`
	LocalRecurso       *bool                     `json:"local_recurso"`
	LitisConsorcio     *bool                     `json:"litis_consorcio"`
	AdvContra          *string                   `json:"adv_contra_nome"`
	DataCad            *string                   `json:"data_cad"`
	DataAltera         *string                   `json:"data_altera"`
	CodProcesso        *string                   `json:"cod_processo"`
	AdvParte           *string                   `json:"adv_parte_nome"`
	TipoProcesso       *string                   `json:"tipo_processo"`
	Justica            *int                      `json:"justica"`
	Tribunal           *int                      `json:"tribunal"`
	NumeroVara         *string                   `json:"numero_vara"`
	Senha              *int                      `json:"senha"`
	Baixado            *int                      `json:"baixado"`
	DataBaixa          *string                   `json:"data_baixa"`
	TipoParte          *int                      `json:"tipo_parte"`
	TipoParteContra    *int                      `json:"tipo_parte_contra"`
	CodAdv             *string                   `json:"cod_adv"`
	DataVerifAnd       *string                   `json:"data_verif_and"`
	Invalido           *int                      `json:"invalido"`
	Judicial           *int                      `json:"judicial"`
	VerifAnd           *int                      `json:"verif_and"`
	Sistema            *int                      `json:"sistema"`
	DataCadastro       *time.Time                `json:"data_cadastro"`
	Excluido           *int                      `json:"excluido"`
	DataExclusao       *string                   `json:"data_exclusao"`
	Assunto            *string                   `json:"assunto"`
	Monitoramento      *int                      `json:"monitoramento"`
	CodParceiro        *int                      `json:"cod_parceiro"`
	ValidoCNJ          *int                      `json:"valido_cnj"`
	ProcessoPartes     []*ProjudProcessoParte    `json:"partes"`
	ProcessoAdvogados  []*ProjudProcessoAdvogado `json:"advogados"`
	processoClientes   []*ProjudProcessoCliente  `json:"clientes"`
}
