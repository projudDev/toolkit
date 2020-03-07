package entities

type Sistema struct {
	Cod            int64
	Descricao      string
	Servidor       string
	IpServidor     string
	Pasta          string
	NomeExecutavel string
	Linguagem      string
	Banco          string
	Status         bool
	Porta          int
	EmailHost      string
	EmailPorta     string
	EmailUser      string
	EmailSenha     string
	EmailTLS       bool
	CreatedAt      time.Time
}
