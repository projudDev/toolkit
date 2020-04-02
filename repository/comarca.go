package repository

import (
	"context"
	"database/sql"

	"errors"
	entities "github.com/projudDev/toolkit/entities"
)

type ProjudComarcaRepo interface {
	FindComarca(ctx context.Context, orgao, tribunal int, numero string) (*entities.ProjudComarca, error)
}

func NewMySQLProjudComarcaRepo(Conn *sql.DB) ProjudComarcaRepo {
	return &mysqlProjudComarcaRepo{Conn: Conn}
}

type mysqlProjudComarcaRepo struct {
	Conn *sql.DB
}

func (this *mysqlProjudComarcaRepo) FindComarca(ctx context.Context, orgao, tribunal int, numero string) (*entities.ProjudComarca, error) {
	projudComarca := new(entities.ProjudComarca)
	query := "Select Cod, Orgao, Tribunal, Numero, Comarca, Local, UF FROM projud_dados.Comarcas where orgao=? and tribunal=? and numero=?;"
	err := this.Conn.QueryRowContext(ctx, query, orgao, tribunal, numero).Scan(
		&projudComarca.ID,
		&projudComarca.Orgao,
		&projudComarca.Tribunal,
		&projudComarca.Numero,
		&projudComarca.Comarca,
		&projudComarca.Local,
		&projudComarca.UF,
	)
	if err != nil && err == sql.ErrNoRows {
		return nil, errors.New("Comarca não encontrada")
	}
	return projudComarca, err
}
