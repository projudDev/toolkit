package digesto

type User struct {
	REF string `json:"$ref,omitempty"`
}

type TribMonitorConfig struct {
	CreateEventOnDiffs   bool `json:"create_event_on_diffs,omitempty"`
	CreateEventOnDistrib bool `json:"create_event_on_distrib,omitempty"`
	CreateEventOnMovs    bool `json:"create_event_on_movs,omitempty"`
	CreateEventOnPubls   bool `json:"create_event_on_publs,omitempty"`
}

type UserCompany struct {
	URI               string             `json:"$uri,omitempty"`
	Name              string             `json:"name,omitempty"`
	DigestoID         string             `json:"digesto_id,omitempty"`
	ProjudID          string             `json:"projud_id,omitempty"`
	APIName           string             `json:"api_name,omitempty"` // Campo livre de até 150 caracteres. Quando fornecido, este texto é enviado em todas as chamadas web-hook desta empresa. Pode ser usado como forma de autenticação.
	CreatedAt         *Date              `json:"created_at,omitempty"`
	EnableModules     []string           `json:"enabled_modules,omitempty"`
	ArchivedAt        *Date              `json:"archived_at,omitempty"` // Se foi excluido (nullable=True). Caso entidate tenha sido excluída, este campo tem o valor da datahora da última exclusão. Caso não esteja excluído, o valor é null. Campo somente-leitura.
	ExpiresAt         *Date              `json:"expires_at,omitempty"`
	ISTrial           bool               `json:"is_trial,omitempty"`
	Comment           string             `json:"comment,omitempty"`
	AdminEmail        string             `json:"admin_email,omitempty"`
	TrialDays         int                `json:"trial_days,omitempty"`
	TribMonitorConfig *TribMonitorConfig `json:"trib_monitor_config,omitempty"` // Dicionário JSON com regras de negócio para monitoramento em tribunais. Ver tabela abaixo. Quando não informado, copiamos o valor da empresa-mãe.
	Users             []*User            `json:"users,omitempty"`
	APIKey            string             `json:"api_key,omitempty"`
}

type MonitoredPerson struct {
	DigestoID           string       `json:"digesto_id,omitempty"`       // ID da pessoa monitorada criada no momento da inserção
	CNPJ                int          `json:"cnpj,omitempty"`             // Usado para encontrar processos desta parte.
	CPF                 int          `json:"cpf,omitempty"`              // Usado para encontrar processos desta parte. Idem, CPF.
	CreatedAT           *Date        `json:"created_at,omitempty"`       // Quando a parte foi criada
	IsActive            bool         `json:"is_active,omitempty"`        //  Se o monitoramento desta parte está ativo.
	ISAdvogado          bool         `json:"is_advogado,omitempty"`      // Se esta pessoa é um advogado
	ISMonitoredTribunal bool         `json:"is_monitored_tribunal"`      // Se o nome é monitorado por distribuições no tribunal (default True)
	ISMonitoredDiario   bool         `json:"is_monitored_diario"`        // Se o nome é monitorado por publicações em diários oficiais (default False)
	ManualRex           bool         `json:"manual_rex,omitempty"`       // Se campos rex e nrex foram editados e nao devem ser reescritos
	Nome                string       `json:"nome,omitempty"`             // Usado para encontrar processos e publicações desta parte. Nome da parte.
	OAB                 string       `json:"oab,omitempty"`              // Usado para encontrar processos e publicações deste advogado. Deve estar no formato UF999999 (duas letras para UF e seis dígitos para número).
	REX                 string       `json:"rex,omitempty"`              // Usado para encontrar processos e publicações desta parte. Expressão regular usada para encontrar processos desta parte.
	NREX                string       `json:"nrex,omitempty"`             // Usado para encontrar processos e publicações desta parte. Expressão regular negativa para encontrar processos desta parte. Ignora estas partes
	AssuntoREX          string       `json:"assunto_rex,omitempty"`      // Usado para encontrar processos desta parte. Considera apenas processos com assunto com essa expressão regular
	NaturezaREX         string       `json:"natureza_rex,omitempty"`     // Usado para encontrar processos desta parte. Considera apenas processos com natureza com essa expressão regular
	RelacaoREX          string       `json:"relacao_rex,omitempty"`      // Usado para encontrar processos desta parte. Considera apenas processos onde parte tem essa relacao no processo
	ComarcaRex          string       `json:"comarca_rex,omitempty"`      // Usado para encontrar processos desta parte. Considera apenas processos em comarca com essa expressão regular
	PartesRex           string       `json:"partes_rex,omitempty"`       // Usado para encontrar processos desta parte. Considera apenas processos onde qualquer uma das partes também dá match nesta expressão regular
	DiariosIDS          []int        `json:"diarios_ids,omitempty"`      // Usado para encontrar publicações desta parte. Serve como filtro. Lista de identificadores de Diários Oficiais que devem ser considerados para recortes em publicações de diários oficiais da justiça.
	TribunaisIDS        []int        `json:"tribunaisIDs,omitempty"`     // Usado para encontrar novos processos distribuídos envolvendo esta parte. Serve como filtro. Lista de identificadores de tribunais que devem ser considerados para monitoramento de distribuições.
	FilterPolo          int          `json:"filter_polo,omitempty"`      // Usado para encontrar processos desta parte. Papel da parte no processo: 0: polo qualquer (default); 1: autor; 2: co-autor 3: réu; 4: neutro
	DistBackDays        int          `json:"dist_back_days,omitempty"`   // Usado para encontrar processos desta parte. Papel da parte no processo: 0: polo qualquer (default); 1: autor; 2: co-autor 3: réu; 4: neutro
	UserCustom          string       `json:"user_custom,omitempty"`      // Usado para encontrar processos desta parte. Papel da parte no processo: 0: polo qualquer (default); 1: autor; 2: co-autor 3: réu; 4: neutro
	DistribuidoFrom     string       `json:"distribuido_from,omitempty"` // Usado para encontrar processos desta parte. Papel da parte no processo: 0: polo qualquer (default); 1: autor; 2: co-autor 3: réu; 4: neutro
	URI                 string       `json:"$uri,omitempty"`
	UserCompany         *UserCompany `json:"user_company,omitempty"`
}

// EndData Representa uma data final
type EndData struct {
	Date int `json:"$date"`
}

// StartData Representa uma data inicial
type StartData struct {
	Date int `json:"$date"`
}

// Date user_company	object	Referência à empresa à qual o pedido pertence. (“UserCompany”). Campo somente-leitura, preenchido automaticamente na criação.
type Date struct {
	Date int64 `json:"$date,omitempty"`
}

// EnvData objeto para envio de datas
type EnvData struct {
	End   EndData   `json:"end_date"`
	Start StartData `json:"start_date"`
}
