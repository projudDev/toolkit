package entities

type ProjudPessoa struct {
	ID            int64      `json:"id"`
	EscritorioID  int64      `json:"escritorio_id"`
	Nome          string     `json:"nome"`
	CPF           *string    `json:"cpf"`
	RG            *string    `json:"rg"`
	RGEmissor     *string    `json:"rg_emissor"`
	RGDataEmissao *string    `json:"rg_data_emissao"`
	Naturalidade  *string    `json:"naturalidade"`
	DataNasc      *string    `json:"data_nasc"`
	Mae           *string    `json:"mae"`
	Pai           *string    `json:"pai"`
	Obs           *string    `json:"obs"`
	DataCad       *string    `json:"data_cad"`
	EstadoCivil   *string    `json:"estado_civil"`
	Tipo          *string    `json:"tipo"`
	IE            *string    `json:"ie"`
	Endereco      *string    `json:"endereco"`
	Numero        *string    `json:"numero"`
	Bairro        *string    `json:"bairro"`
	CEP           *string    `json:"cep"`
	Cidade        *string    `json:"cidade"`
	UF            *string    `json:"uf"`
	Sexo          *string    `json:"sexo"`
	Contato       *string    `json:"contato"`
	RamoAtividade *string    `json:"ramo_atividade"`
	Fone2         *string    `json:"fone2"`
	Fone3         *string    `json:"fone3"`
	Fone4         *string    `json:"fone4"`
	Fone          *string    `json:"fone"`
	Email         *string    `json:"email"`
	Site          *string    `json:"site"`
	Complemento   *string    `json:"complemento"`
	Profissao     *string    `json:"profissao"`
	Classificacao *string    `json:"classificacao"`
	DataAltera    *string    `json:"data_altera"`
	CodAnt        *int       `json:"cod_ant"`
	DataCadastro  *time.Time `json:"data_cadastro"`
	Excluido      *int       `json:"escluido"`
	DataExcludao  *time.Time `json:"data_exclusao"`
	EscavadorID   *int       `json:"escavador_id"`
}
