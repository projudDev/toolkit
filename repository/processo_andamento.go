package repository

import (
	"context"
	"database/sql"
	"time"

	entities "github.com/projudDev/toolkit/entities"
)

type ProjudProcessoAndamentoRepo interface {
	Create(ctx context.Context, processoAndamento *entities.ProjudProcessoAndamento) (int64, error)
	IsExist(ctx context.Context, processoID int64, data time.Time, descricao string) (existe bool, err error)
}

func NewMySQLProjudProcessoAndamentoRepo(Conn *sql.DB) ProjudProcessoAndamentoRepo {
	return &mysqlProjudProcessoAndamentoRepo{Conn: Conn}
}

type mysqlProjudProcessoAndamentoRepo struct {
	Conn *sql.DB
}

func (this *mysqlProjudProcessoAndamentoRepo) IsExist(ctx context.Context, processoID int64, data time.Time, descricao string) (existe bool, err error) {
	query := "SELECT IF(COUNT(*),'true','false') FROM projud_andamentos.processosandamentos WHERE codproc=? and data=? and descricao=?;"
	row := this.Conn.QueryRowContext(ctx, query, processoID, data.Format("2006-01-02"), descricao)
	err = row.Scan(&existe)
	return
}

func (this *mysqlProjudProcessoAndamentoRepo) Create(ctx context.Context, processoAndamento *entities.ProjudProcessoAndamento) (int64, error) {
	query := "INSERT INTO projud_andamentos.processosandamentos(codproc, descricao, data, datacad) values (?,?,?,?);"
	stmt, err := this.Conn.PrepareContext(ctx, query)
	if err != nil {
		return -1, err
	}
	res, err := stmt.ExecContext(ctx,
		&processoAndamento.ProcessoID,
		&processoAndamento.Descricao,
		&processoAndamento.Data,
		&processoAndamento.DataCad,
	)
	if err != nil {
		return -1, err
	}

	defer stmt.Close()

	return res.LastInsertId()
}
