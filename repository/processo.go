package repository

import (
	"context"
	"database/sql"

	entities "github.com/projudDev/toolkit/entities"
)

type ProjudProcessoRepo interface {
	Create(ctx context.Context, processo *entities.ProjudProcesso) (int64, error)
}

func NewMySQLProjudProcessoRepo(Conn *sql.DB) ProjudProcessoRepo {
	return &mysqlProjudProcessoRepo{Conn: Conn}
}

type mysqlProjudProcessoRepo struct {
	Conn *sql.DB
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
