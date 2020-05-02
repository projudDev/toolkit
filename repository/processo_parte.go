package repository

import (
	"context"
	"database/sql"

	entities "github.com/projudDev/toolkit/entities"
)

type ProjudProcessoParteRepo interface {
	Create(ctx context.Context, processoParte *entities.ProjudProcessoParte) (int64, error)
	IsExist(ctx context.Context, processoID, clienteID int64) (existe bool, err error)
}

func NewMySQLProjudProcessoParteRepo(Conn *sql.DB) ProjudProcessoParteRepo {
	return &mysqlProjudProcessoParteRepo{Conn: Conn}
}

type mysqlProjudProcessoParteRepo struct {
	Conn *sql.DB
}

func (this *mysqlProjudProcessoParteRepo) IsExist(ctx context.Context, processoID, clienteID int64) (existe bool, err error) {
	query := "SELECT IF(COUNT(*),'true','false') FROM projud_dados.processospartes WHERE codcli=? and codproc=?;"
	row := this.Conn.QueryRowContext(ctx, query, clienteID, processoID)
	err = row.Scan(&existe)
	return
}

func (this *mysqlProjudProcessoParteRepo) Create(ctx context.Context, processoParte *entities.ProjudProcessoParte) (int64, error) {
	query := "INSERT INTO projud_dados.processospartes(codcli, codproc, polo) values (?,?,?);"
	stmt, err := this.Conn.PrepareContext(ctx, query)
	if err != nil {
		return -1, err
	}
	res, err := stmt.ExecContext(ctx,
		processoParte.ClienteID,
		processoParte.ProcessoID,
		processoParte.Polo,
	)
	if err != nil {
		return -1, err
	}

	defer stmt.Close()

	return res.LastInsertId()
}
