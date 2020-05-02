package repository

import (
	"context"
	"database/sql"

	entities "github.com/projudDev/toolkit/entities"
)

type ProjudAgendaRepo interface {
	Create(ctx context.Context, projudAgenda *entities.ProjudAgenda) (int64, error)
}

func NewMySQLProjudAgendaRepo(Conn *sql.DB) ProjudAgendaRepo {
	return &mysqlProjudAgendaRepo{Conn: Conn}
}

type mysqlProjudAgendaRepo struct {
	Conn *sql.DB
}

func (this *mysqlProjudAgendaRepo) Create(ctx context.Context, projudAgenda *entities.ProjudAgenda) (int64, error) {
	query := "INSERT INTO projud_dados.agenda(codesc, data, horai, horaf, diatodo, evento, obs, prazo, concluido, vinculo, codvinculo, dataf, coduso, tipoevento, codproc, codcli, local, subcategoria, cor, incluir_telefone, excluido, codpub, privacidade, id_fuso_horario) values (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?);"
	stmt, err := this.Conn.PrepareContext(ctx, query)
	if err != nil {
		return -1, err
	}
	res, err := stmt.ExecContext(ctx,
		projudAgenda.EscritorioID,
		projudAgenda.Data.Format("2006-01-02"),
		projudAgenda.HoraInicial,
		projudAgenda.HoraFinal,
		projudAgenda.DiaTodo,
		projudAgenda.Evento,
		projudAgenda.Obs,
		projudAgenda.Prazo,
		projudAgenda.Cloncluido,
		projudAgenda.Vinculo,
		projudAgenda.CodVinculo,
		projudAgenda.DataFinal,
		projudAgenda.CodUso,
		projudAgenda.TipoEvento,
		projudAgenda.ProcessoID,
		projudAgenda.ClienteID,
		projudAgenda.Local,
		projudAgenda.SubCategoria,
		projudAgenda.Cor,
		projudAgenda.IncluirTelefone,
		projudAgenda.Excluido,
		projudAgenda.CodPub,
		projudAgenda.Privacidade,
		projudAgenda.FusoHorarioID,
	)
	if err != nil {
		return -1, err
	}

	defer stmt.Close()

	return res.LastInsertId()
}
