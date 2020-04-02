package repository

import (
	"context"
	"database/sql"

	"errors"
	entities "github.com/projudDev/toolkit/entities"
)

type ProjudPessoaRepo interface {
	Create(ctx context.Context, projudPessoa *entities.ProjudPessoa) (int64, error)
	FindByNome(ctx context.Context, EscritorioID int64, nome string) (*entities.ProjudPessoa, error)
	GetMaxID(ctx context.Context) (ID int64, err error)
}

func NewMySQLProjudPessoaRepo(Conn *sql.DB) ProjudPessoaRepo {
	return &mysqlProjudPessoaRepo{Conn: Conn}
}

type mysqlProjudPessoaRepo struct {
	Conn *sql.DB
}

func (this *mysqlProjudPessoaRepo) GetMaxID(ctx context.Context) (ID int64, err error) {
	query := "SELECT COALESCE(MAX(Cod)+1, 1) as Cod FROM clientes"
	row := this.Conn.QueryRowContext(ctx, query)
	err = row.Scan(&ID)
	return
}

func (this *mysqlProjudPessoaRepo) FindByNome(ctx context.Context, EscritorioID int64, nome string) (*entities.ProjudPessoa, error) {
	projudPessoa := new(entities.ProjudPessoa)
	query := "SELECT cod, codesc, nome FROM clientes WHERE codesc=? and nome=?;"
	err := this.Conn.QueryRowContext(ctx, query, EscritorioID, nome).Scan(
		&projudPessoa.ID,
		&projudPessoa.EscritorioID,
		&projudPessoa.Nome,
	)
	if err != nil && err == sql.ErrNoRows {
		return nil, errors.New("Pessoa não encontrada")
	}
	return projudPessoa, err
}

func (this *mysqlProjudPessoaRepo) Create(ctx context.Context, projudPessoa *entities.ProjudPessoa) (int64, error) {
	query := "INSERT INTO clientes(codesc, nome, cpf, rg, emissor, dataemissao, naturalidade, datanasc, mae, pai, obs, datacad, estadocivil, tipo, ie, endereco, numero, bairro, cep, cidade, uf, sexo, contato, ramoatividade, fone2, fone3, fone4, fone, email, site, complemento, profissao, classificacao, dataaltera, codant, datacadastro, excluido, dataexclusao, idescavador) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"

	stmt, err := this.Conn.PrepareContext(ctx, query)
	if err != nil {
		return -1, err
	}
	res, err := stmt.ExecContext(ctx,
		projudPessoa.EscritorioID,
		projudPessoa.Nome,
		projudPessoa.CPF,
		projudPessoa.RG,
		projudPessoa.RGEmissor,
		projudPessoa.RGDataEmissao,
		projudPessoa.Naturalidade,
		projudPessoa.DataNasc,
		projudPessoa.Mae,
		projudPessoa.Pai,
		projudPessoa.Obs,
		projudPessoa.DataCad,
		projudPessoa.EstadoCivil,
		projudPessoa.Tipo,
		projudPessoa.IE,
		projudPessoa.Endereco,
		projudPessoa.Numero,
		projudPessoa.Bairro,
		projudPessoa.CEP,
		projudPessoa.Cidade,
		projudPessoa.UF,
		projudPessoa.Sexo,
		projudPessoa.Contato,
		projudPessoa.RamoAtividade,
		projudPessoa.Fone2,
		projudPessoa.Fone3,
		projudPessoa.Fone4,
		projudPessoa.Fone,
		projudPessoa.Email,
		projudPessoa.Site,
		projudPessoa.Complemento,
		projudPessoa.Profissao,
		projudPessoa.Classificacao,
		projudPessoa.DataAltera,
		projudPessoa.CodAnt,
		projudPessoa.DataCadastro,
		projudPessoa.Excluido,
		projudPessoa.DataExcludao,
		projudPessoa.EscavadorID,
	)
	if err != nil {
		return -1, err
	}

	defer stmt.Close()
	return res.LastInsertId()
}
