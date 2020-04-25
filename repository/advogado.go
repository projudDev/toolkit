package repository

import (
	"context"
	"database/sql"

	"errors"
	entities "github.com/projudDev/toolkit/entities"
)

type ProjudAdvogadoRepo interface {
	FindByNome(ctx context.Context, EscritorioID int64, nome string) (*entities.ProjudAdvogado, error)
}

func NewMySQLProjudAdvogadoRepo(Conn *sql.DB) ProjudAdvogadoRepo {
	return &mysqlProjudAdvogadoRepo{Conn: Conn}
}

type mysqlProjudAdvogadoRepo struct {
	Conn *sql.DB
}

func (this *mysqlProjudAdvogadoRepo) FindByNome(ctx context.Context, EscritorioID int64, nome string) (*entities.ProjudAdvogado, error) {
	projudAdvogado := new(entities.ProjudAdvogado)
	query := "SELECT cod, codesc, nome FROM projud_dados.advogados WHERE codesc=? and nome=?;"
	err := this.Conn.QueryRowContext(ctx, query, EscritorioID, nome).Scan(
		&projudAdvogado.ID,
		&projudAdvogado.EscritorioID,
		&projudAdvogado.Nome,
	)
	if err != nil && err == sql.ErrNoRows {
		return nil, errors.New("Advogado não encontrado")
	}
	return projudAdvogado, err
}
