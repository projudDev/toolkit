package repository

import (
	"context"
	"database/sql"

	entities "github.com/projudDev/toolkit/entities"
)

type ProjudNotificacaoRepo interface {
	Create(ctx context.Context, projudNotificacao *entities.ProjudNotificacao) (int64, error)
}

func NewMySQLProjudNotificacaoRepo(Conn *sql.DB) ProjudNotificacaoRepo {
	return &mysqlProjudNotificacaoRepo{Conn: Conn}
}

type mysqlProjudNotificacaoRepo struct {
	Conn *sql.DB
}

func (this *mysqlProjudNotificacaoRepo) Create(ctx context.Context, projudNotificacao *entities.ProjudNotificacao) (int64, error) {
	query := "INSERT INTO projud_dados.notificacoes(coduso, codesc, titulo, texto, link, lido, tipo) values (?,?,?,?,?,?,?);"

	stmt, err := this.Conn.PrepareContext(ctx, query)
	if err != nil {
		return -1, err
	}
	res, err := stmt.ExecContext(ctx,
		&projudNotificacao.UsuarioID,
		&projudNotificacao.EscritorioID,
		&projudNotificacao.Titulo,
		&projudNotificacao.Texto,
		&projudNotificacao.Link,
		&projudNotificacao.Lido,
		&projudNotificacao.Tipo,
	)
	if err != nil {
		return -1, err
	}

	defer stmt.Close()
	return res.LastInsertId()
}
