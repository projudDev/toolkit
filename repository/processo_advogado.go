package repository

import (
	"context"
	"database/sql"

	entities "github.com/projudDev/toolkit/entities"
)

type ProjudProcessoAdvogadoRepo interface {
	Create(ctx context.Context, processoAdvogado *entities.ProjudProcessoAdvogado) (int64, error)
}

func NewMySQLProjudProcessoAdvogadoRepo(Conn *sql.DB) ProjudProcessoAdvogadoRepo {
	return &mysqlProjudProcessoAdvogadpoRepo{Conn: Conn}
}

type mysqlProjudProcessoAdvogadpoRepo struct {
	Conn *sql.DB
}

func (this *mysqlProjudProcessoAdvogadpoRepo) Create(ctx context.Context, processoAdvogado *entities.ProjudProcessoAdvogado) (int64, error) {
	query := "INSERT INTO projud_dados.processospartes(codadv, codproc, polo) values (?,?,?);"

	stmt, err := this.Conn.PrepareContext(ctx, query)
	if err != nil {
		return -1, err
	}
	res, err := stmt.ExecContext(ctx,
		processoAdvogado.AdvogadoID,
		processoAdvogado.ProcessoID,
		processoAdvogado.Polo,
	)
	if err != nil {
		return -1, err
	}

	defer stmt.Close()
	return res.LastInsertId()
}
