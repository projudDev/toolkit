package repository

import (
	"context"
	"database/sql"
	entities "github.com/projudDev/toolkit/entities"
)

type ProjudUsuarioRepo interface {
	FindByEscritorioID(ctx context.Context, id int64) ([]*entities.ProjudUsuario, error)
}

func NewMySQLProjudUsuarioRepo(Conn *sql.DB) ProjudUsuarioRepo {
	return &mysqlProjudUsuarioRepo{Conn: Conn}
}

type mysqlProjudUsuarioRepo struct {
	Conn *sql.DB
}

func (this *mysqlProjudUsuarioRepo) FindByEscritorioID(ctx context.Context, escritorioID int64) ([]*entities.ProjudUsuario, error) {
	query := `SELECT cod, codesc, nome, login, senha, automatico, tipo, ultimoacesso, ultimoacessodata, email, ativo, corc,corm,fone,notificasms, notificaand, notificapub, acessos, datacadastro, agendagoogle, excluido, dataexclusao, aob, aob_uf, oab_letra, imagem, updated_at FROM projud_dados.login WHERE codesc=?`
	return this.fetch(ctx, query, escritorioID)
}

func (this *mysqlProjudUsuarioRepo) fetch(ctx context.Context, query string, args ...interface{}) ([]*entities.ProjudUsuario, error) {
	rows, err := this.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	payload := make([]*entities.ProjudUsuario, 0)
	for rows.Next() {
		projudUsuario := new(entities.ProjudUsuario)
		err := rows.Scan(
			&projudUsuario.ID,
			&projudUsuario.EscritorioID,
			&projudUsuario.Nome,
			&projudUsuario.Login,
			&projudUsuario.Senha,
			&projudUsuario.Automatico,
			&projudUsuario.Tipo,
			&projudUsuario.UltimoAcesso,
			&projudUsuario.UltimoAcessoData,
			&projudUsuario.Email,
			&projudUsuario.Ativo,
			&projudUsuario.CorC,
			&projudUsuario.CorM,
			&projudUsuario.Fone,
			&projudUsuario.NotificaSMS,
			&projudUsuario.NotificaAnd,
			&projudUsuario.NotificaPub,
			&projudUsuario.Acessos,
			&projudUsuario.DataCadastro,
			&projudUsuario.AgendaGoogle,
			&projudUsuario.Excluido,
			&projudUsuario.DataExclusao,
			&projudUsuario.OAB,
			&projudUsuario.OABUF,
			&projudUsuario.OABLetra,
			&projudUsuario.Imagem,
			&projudUsuario.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		payload = append(payload, projudUsuario)
	}
	return payload, nil
}
