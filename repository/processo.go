package repository

import (
	"context"
	"database/sql"
	"errors"

	entities "github.com/projudDev/toolkit/entities"
)

type ProjudProcessoRepo interface {
	Create(ctx context.Context, processo *entities.ProjudProcesso) (int64, error)
	IsExist(ctx context.Context, escritorioID int64, autos string) (existe bool, err error)
	SetClienteID(ctx context.Context, processoID, clienteID int64) error
	GetMaxID(ctx context.Context) (ID int64, err error)
	SetClienteNome(ctx context.Context, processoID int64, nome string) error
}

func NewMySQLProjudProcessoRepo(Conn *sql.DB) ProjudProcessoRepo {
	return &mysqlProjudProcessoRepo{Conn: Conn}
}

type mysqlProjudProcessoRepo struct {
	Conn *sql.DB
}

func (this *mysqlProjudProcessoRepo) GetMaxID(ctx context.Context) (ID int64, err error) {
	query := "SELECT COALESCE(MAX(Cod)+1, 1) as Cod FROM processos"
	row := this.Conn.QueryRowContext(ctx, query)
	err = row.Scan(&ID)
	return
}

func (this *mysqlProjudProcessoRepo) IsExist(ctx context.Context, escritorioID int64, autos string) (existe bool, err error) {
	query := "SELECT IF(COUNT(*),'true','false') FROM projud_dados.processos WHERE codesc=? and autos=?"
	row := this.Conn.QueryRowContext(ctx, query, escritorioID, autos)
	err = row.Scan(&existe)
	return
}

func (this *mysqlProjudProcessoRepo) Create(ctx context.Context, processo *entities.ProjudProcesso) (int64, error) {

	query := "INSERT INTO projud_dados.processos(codesc, codcli, autos, nome, comarca, vara, acao, protocolo, valorcausa, honorarios, partecontra,situacao, obs, ativo, motivoencerramento, recurso, localrecurso, litisconsorcio, advcontra, datacad, dataaltera, codprocesso, advparte, tipoprocesso, justica, tribunal, numerovara, senha, baixado, databaixa, tipoparte, tipopartecontra, codadv, dataverifand, invalido, judicial, verifand, sistema, datacadastro, excluido, dataexclusao, assunto, monitoramento, codparceiro, validocnj) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"

	stmt, err := this.Conn.PrepareContext(ctx, query)
	if err != nil {
		return -1, err
	}
	res, err := stmt.ExecContext(ctx,
		processo.EscritorioID,
		processo.ClienteID,
		processo.Autos,
		processo.Nome,
		processo.Comarca,
		processo.Vara,
		processo.Acao,
		processo.Protocolo,
		processo.ValorCausa,
		processo.Honorarios,
		processo.ParteContra,
		processo.Situacao,
		processo.Obs,
		processo.Ativo,
		processo.MotivoEncerramento,
		processo.Recurso,
		processo.LocalRecurso,
		processo.LitisConsorcio,
		processo.AdvContra,
		processo.DataCad,
		processo.DataAltera,
		processo.CodProcesso,
		processo.AdvParte,
		processo.TipoProcesso,
		processo.Justica,
		processo.Tribunal,
		processo.NumeroVara,
		processo.Senha,
		processo.Baixado,
		processo.DataBaixa,
		processo.TipoParte,
		processo.TipoParteContra,
		processo.CodAdv,
		processo.DataVerifAnd,
		processo.Invalido,
		processo.Judicial,
		processo.VerifAnd,
		processo.Sistema,
		processo.DataCadastro,
		processo.Excluido,
		processo.DataExclusao,
		processo.Assunto,
		processo.Monitoramento,
		processo.CodParceiro,
		processo.ValidoCNJ,
	)
	if err != nil {
		return -1, err
	}

	defer stmt.Close()
	return res.LastInsertId()
}

func (this *mysqlProjudProcessoRepo) SetClienteID(ctx context.Context, processoID, clienteID int64) error {
	query := "UPDATE processos SET codcli=? WHERE cod=?"
	stmt, err := this.Conn.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	res, err := stmt.ExecContext(ctx, clienteID, processoID)
	if err != nil {
		return err
	}
	if numRols, _ := res.RowsAffected(); numRols == 0 {
		return errors.New("Nenhum registro afetado, verifique o cod")
	}
	return nil
}

func (this *mysqlProjudProcessoRepo) SetClienteNome(ctx context.Context, processoID int64, nome string) error {
	query := "UPDATE processos SET nome=? WHERE cod=?"
	stmt, err := this.Conn.PrepareContext(ctx, query)
	if err != nil {
		return err
	}
	res, err := stmt.ExecContext(ctx, nome, processoID)
	if err != nil {
		return err
	}
	if numRols, _ := res.RowsAffected(); numRols == 0 {
		return errors.New("Nenhum registro afetado, verifique o cod")
	}
	return nil
}
