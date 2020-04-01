package repository

import (
	"context"
	"database/sql"

	entities "github.com/projudDev/toolkit/entities"
)

type ProjudProcessoParteRepo interface {
	Create(ctx context.Context, processoParte *entities.ProjudProcessoParte) (int64, error)
}

func NewMySQLProjudProcessoParteRepo(Conn *sql.DB) ProjudProcessoParteRepo {
	return &mysqlProjudProcessoParteRepo{Conn: Conn}
}

type mysqlProjudProcessoParteRepo struct {
	Conn *sql.DB
}

func (this *mysqlProjudProcessoParteRepo) Create(ctx context.Context, processoParte *entities.ProjudProcessoParte) (int64, error) {
	query := "INSERT INTO processospartes(codcli, codproc, polo) values (?,?,?);"

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
