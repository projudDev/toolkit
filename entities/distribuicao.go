package entities

import (
	"time"
)

type DistribuicaoEvento struct {
	ID               int64
	DigestoID        int64                  `json:"id"`
	TargetURL        string                 `json:"target_url"`
	SourceUserCustom interface{}            `json:"source_user_custom"`
	CreatedAt        string                 `json:"created_at"`
	SourceURL        []string               `json:"source_url"`
	TargetNumber     string                 `json:"target_number"`
	EvtType          int                    `json:"evt_type"`
	Data             []DistribuicaoProcesso `json:"data"`
}

type DistribuicaoSourceURL struct {
	ID       int64  `json:"id"`
	EventoID int64  `json:"evento_id"`
	Url      string `json:"url"`
}

type DistribuicaoProcesso struct {
	ID                 int64              `json:"id"`
	DigestoPersonID    int64              `json:"person_id"`
	EventoID           int64              `json:"evento_id"`
	NumeroAlternativo  string             `json:"numeroAlternativo"`
	Anexos             [][]interface{}    `json:"anexos"`
	Publicacoes        []interface{}      `json:"publicacoes"`
	Vara               string             `json:"vara"`
	VaraOriginal       string             `json:"vara_original"`
	ClasseNatureza     string             `json:"classeNatureza"`
	Movs               [][]interface{}    `json:"movs"`
	DistribuicaoTipo   string             `json:"distribuicaoTipo"`
	SituacaoSituacaoID *int               `json:"situacao_situacaoID"`
	Area               string             `json:"area"`
	AssuntoExtra       *string            `json:"assuntoExtra"`
	Audiencias         [][]string         `json:"audiencias"`
	Extinto            int                `json:"extinto"`
	SentencaData       *time.Time         `json:"sentencaData"`
	AlteradoEm         string             `json:"alteradoEm"`
	Foro               string             `json:"foro"`
	Situacao           *string            `json:"situacao"`
	Valor              *float64           `json:"valor"`
	Numero             string             `json:"numero"`
	Customs            DistribuicaoCustom `json:"customs"`
	Comarca            string             `json:"comarca"`
	TribunalID         int                `json:"tribunalID"`
	Tribunal           string             `json:"tribunal"`
	ProcessoID         int                `json:"processoID"`
	DistribuicaoData   string             `json:"distribuicaoData"`
	Arquivado          bool               `json:"arquivado"`
	Classes            []string           `json:"classes"`
	Acessos            string             `json:"acessos"`
	Uf                 string             `json:"uf"`
	CriadoEm           string             `json:"criadoEm"`
	Partes             [][]interface{}    `json:"partes"`
	DistribuicaoPartes []DistribuicaoParte
}


type DistribuicaoCustom struct {
	ID                  int64  `json:"id"`
	ProcessoID      int64  `json:"processo_id"`
	SourceAlgoAmplitude string `json:"source_algo_amplitude"`
	SourceAlgoDias      string `json:"source_algo_dias"`
	SourceProcSource    string `json:"source_proc_source"`
	SourceAlgoScan      string `json:"source_algo_scan"`
}

/*
0 advogado ID (int),
1 nome (texto),
2 oab (texto),
3 CPF (texto),
4 UF (texto)
14 CEP da parte (int). Quando for possível extrair da inicial. [versão API 2],
15 ID da entidade monitored_person correspondente a esta parte. Útil quando este processo for enviado como uma nova distribuição de uma parte monitorda e você deseja saber qual parte monitorada de uma empresa deu match para esse processo. (int) [versão API 2]
16 parte é pessoa física (boolean) [versão API 3]*/
type DistribuicaoParte struct {
	Id                   int64 //0 ID interno da participação dessa parte nesse processo (int),
	Id_distribuicao      int64
	Id_interno           float64 //1 ID interno para essa parte (int),
	Nome                 string  //2 nome da parte (texto),
	Nome_normalizado     string  //3 nome da parte normalizado: sem sufixos de empresas, pontuações etc (texto),
	Cnpj                 float64 //4 cnpj: quando disponibilizado por algum tribunal ou na inicial de processos digitais (int),
	CnpjStr              string
	Cpf                  interface{}     //5 cpf: quando disponibilizado por algum tribunal ou na inicial de processos digitais (int),
	Documento            interface{}     //6 documento: CNPJ ou CPF formatado, dependendo da natureza da parte (texto),
	ParteRelacaoID       float64         //7 parteRelacaoID: ID do tipo de relação/papel da parte no processo (int),
	RelacaoNormalizado   string          //8 relacaoNormalizado: nome da relação/papel que esta parte desempenha no processo (texto),
	Advogados            [][]interface{} //9 advogados: lista de tuplas, cada tupla representa um advogado desta parte e é composta pelos valores:
	Autora               bool            //10 parte é autora (booleano). Quando não for possível fazer a classificação do papel da parte, estes quatro indicadores virão como false,
	Co_autora            bool            //11 parte é co-autora (booleano),
	Re                   bool            //12 parte é ré (booleano),
	Neutra               bool            //13 parte é neutra (booleano),
	DistribuicaoAdvogado []DistribuicaoAdvogado
}

type DistribuicaoAdvogado struct {
	Id       int64
	Id_parte int64
	Id_adv   int    //0 advogado ID (int),
	Nome     string //1 nome (texto),
	Oab      string //2 oab (texto),
	Cpf      string //3 CPF (texto),
	Uf       string //4 UF (texto)
}

type DistribuicaoAnexo struct {
	Id              int64 //0 processo Anexo ID,
	Id_distribuicao int64
	Digesto_id      int64
	Url             string      //1 endereco HTTP para download do anexo (texto),
	Tipo_anexo      int         //2 tipo de Anexo (inteiro): 1 - Inicial, 2 - Sentenca, 3 - Outros, 4 - Ajuizamento
	Data_publicacao interface{} //3 data de publicacao (YYYY-MM-DDTHH:MM:ss),
	Conteudo        string      //4 conteudo em modo texto (texto) Atenção: nos detalhes dos processos enviados numa distribuição (evt_type=4), este campo é omitido/pulado e nesta posição 4 é enviado a ``data de obtenção`` e assim por diante, ou seja, a tupla de anexos fica com 7 elementos.
	Data_obtencao   interface{} //5 data de obtenção (YYYY-MM-DDTHH:MM:ss),
	Id_movimentacao interface{} //6 ID movimentação: quando o tribunal associa o anexo a uma movimentação, indicamos o ID interno Digesto da mesma (5o. campo das tuplas em movs)
	Titulo_anexo    string      //7 título do anexo, quando disponível (texto)
}

type DistribuicaoAudiencia struct {
	Id        int64     //id do banco
	Data_hora time.Time //0 data e hora (YYYY-MM-DD HH:MM:SS). Quando o horário não for informado pelo tribunal, enviamos o horário 00:00.
	Local     string    //1 local (texto),
	Tipo      string    //2 tipo (texto),
	Situacao  string    //3 situacao (texto) (enviado apenas quando o parâmetro get_situacao_audiencia=true é passado na chamada)
}

type DistibuicaoMovimentacao struct {
	Id               int64
	Id_distribuicao  int64
	Data             string          //0 data (texto, com a data no formato YYYY-MM-DD), a data da movimentação indicada pelo tribunal.
	Tipo             string          //1 tipo (texto, até 255 caracteres), o tipo da movimentação é exatamente o informado pelo tribunal como Titulo da movimentacao. Nem todos disponibilizam um titulo. As vezes é só uma abreviação. Há bastante variação entre os Tribunais.
	Texto            string          //2 texto (texto, tamanho variável).
	Nome_juiz        string          //3 nome do juiz (texto, até 255 caracteres).
	Id_interno       string          //4 ID interno da movimentação (texto, até 16 caracteres). Na [versão API 1,2,3] é um (int).
	Tipo_inteligente [][]interface{} //5 tipo inteligente (array de tuplas), é - quando disponível - a classificação automática pelo Digesto da movimentação. Cada tupla dessa lista representa um assunto mencionado na movimentação. Quando não disponível, o valor é uma lista vazia. Veja Tipos padronizados Digesto para andamentos e publicações processuais com os valores que cada tupla pode assumir.
}

type DistribuicaoClasse struct {
	Id              int64
	Id_distribuicao int64
	Classe          string
}

type AutoGenerated []struct {
	APIName          string    `json:"api_name"`
	TargetURL        string    `json:"target_url"`
	SourceUserCustom *string   `json:"source_user_custom"`
	TargetNumber     string    `json:"target_number"`
	EvtType          int       `json:"evt_type"`
	CreatedAt        time.Time `json:"created_at"`
	Data             []struct {
		RegionalCnj           bool            `json:"regional_cnj"`
		NumeroAlternativo     *string         `json:"numeroAlternativo"`
		Anexos                []interface{}   `json:"anexos"`
		Vara                  string          `json:"vara"`
		Partes                [][]interface{} `json:"partes"`
		Movs                  [][]interface{} `json:"movs"`
		DistribuicaoTipo      string          `json:"distribuicaoTipo"`
		SituacaoSituacaoID    int             `json:"situacao_situacaoID"`
		Juiz                  *string         `json:"juiz"`
		Area                  string          `json:"area"`
		AssuntoExtra          *string         `json:"assuntoExtra"`
		Liminar               *bool           `json:"liminar"`
		Audiencias            interface{}     `json:"audiencias"`
		VaraOriginal          string          `json:"vara_original"`
		Extinto               int             `json:"extinto"`
		Gratuita              interface{}     `json:"gratuita"`
		Valor                 interface{}     `json:"valor"`
		AlteradoEm            string          `json:"alteradoEm"`
		FonteSistema          string          `json:"fonte_sistema"`
		Foro                  string          `json:"foro"`
		Situacao              string          `json:"situacao"`
		Instancia             int             `json:"instancia"`
		ProcessosRelacionados []interface{}   `json:"processosRelacionados"`
		Numero                string          `json:"numero"`
		Flag                  int             `json:"flag"`
		ComarcaCnj            string          `json:"comarca_cnj"`
		Comarca               string          `json:"comarca"`
		TribunalID            int             `json:"tribunalID"`
		ForoCnj               string          `json:"foro_cnj"`
		Tribunal              string          `json:"tribunal"`
		ProcessoID            int             `json:"processoID"`
		DistribuicaoData      string          `json:"distribuicaoData"`
		SentencaData          interface{}     `json:"sentencaData"`
		Arquivado             bool            `json:"arquivado"`
		Classes               []interface{}   `json:"classes"`
		Acessos               string          `json:"acessos"`
		Uf                    string          `json:"uf"`
		CriadoEm              string          `json:"criadoEm"`
		ClasseNatureza        string          `json:"classeNatureza"`
	} `json:"data"`
	ID        int      `json:"id"`
	SourceURL []string `json:"source_url"`
}
