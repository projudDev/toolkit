package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/projudDev/toolkit/driver"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

func main() {
	conn, err := driver.NewConnectionMysql()
	if err != nil {
		log.Fatalln(err.Error())
	}
	advogados := getAdvogados(conn.SQL)
	for _, adv := range advogados {
		importarNome(conn.SQL, adv.Cod, adv.CodEsc)
	}

}

type Advogado struct {
	Cod    int
	CodEsc int
}

func getAdvogados(db *sql.DB) []Advogado {
	var payload []Advogado
	rows, err := db.Query("SELECT cod FROM `escritorios` where cod > 1001002")
	if err != nil {
		log.Println(err.Error())
		return payload
	}
	defer rows.Close()

	for rows.Next() {
		var advogado Advogado
		err = rows.Scan(&advogado.CodEsc)
		if err != nil {
			log.Println(err.Error())
			return payload
		}
		payload = append(payload, advogado)
	}
	return payload
}

func GetEscritorio(codesc int) (Escritorio, error) {
	var escritorio Escritorio
	url := "http://api.projud.com.br:8097/api/v1/listar/escritorio/" + strconv.Itoa(codesc)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return escritorio, err
	}
	var bearer = "Bearer " + "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c3VhcmlvIjp7ImlkIjo1LCJlbXByZXNhIjoiUFJPSlVEIiwidG9rZW4iOiI2N2VkOGJjOGFlZTkxN2E5MzI5ZTU3ZTZjOWNmZDM1ZSJ9LCJpc3MiOiJQUk9KVUQifQ.qB40JDWnAKzXLQZJ41UNq0sa8pSxIJJ4Rum3eoBjaofMg8-x6-F83e08UqebNGKTigv1zyb1-s0hxFUKko8xW1fgPu_2uf6PNzq1EME6UVt3gRKT0s1oW-pKMdE6YstimRn_SXfK3gx-y_ruUQlHfFHu_xqGB7YuG9zM2WHB0Ns"

	req.Header.Add("Authorization", bearer)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return escritorio, err
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return escritorio, err
	}
	if res.StatusCode != 200 {
		return escritorio, errors.New("Error: " + string(b))
	}
	err = json.Unmarshal(b, &escritorio)
	if err != nil {
		return escritorio, err
	}
	return escritorio, nil

}
func importarNome(db *sql.DB, cod int, codesc int) {
	fmt.Println("# Importando Advogado:", cod, "Escritório:", codesc)
	escritorio, err := GetEscritorio(codesc)
	if err != nil {
		log.Println(err.Error())
		return
	}
	url := "http://api.projud.com.br:8097/api/v1/listar/nomespesquisa/escritorio/" + strconv.Itoa(codesc)

	var bearer = "Bearer " + "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c3VhcmlvIjp7ImlkIjo1LCJlbXByZXNhIjoiUFJPSlVEIiwidG9rZW4iOiI2N2VkOGJjOGFlZTkxN2E5MzI5ZTU3ZTZjOWNmZDM1ZSJ9LCJpc3MiOiJQUk9KVUQifQ.qB40JDWnAKzXLQZJ41UNq0sa8pSxIJJ4Rum3eoBjaofMg8-x6-F83e08UqebNGKTigv1zyb1-s0hxFUKko8xW1fgPu_2uf6PNzq1EME6UVt3gRKT0s1oW-pKMdE6YstimRn_SXfK3gx-y_ruUQlHfFHu_xqGB7YuG9zM2WHB0Ns"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err.Error())
		return
	}
	req.Header.Add("Authorization", bearer)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Println(err.Error())
		return
	}
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err.Error())
		return
	}

	if res.StatusCode != 200 {
		log.Println("importarNome.Error:", string(b))
		return
	}

	var nome Nomes
	err = json.Unmarshal(b, &nome)
	if err != nil {
		log.Println(err.Error())
		return
	}

	for _, item := range nome.Item {

		var termoMonitorado TermoMonitorado

		termoMonitorado.EscID = item.CodEscritorio
		termoMonitorado.Termo = item.Nome
		termoMonitorado.SolucionareID = item.CodNome
		termoMonitorado.Status = 1

		if escritorio.Bloqueado {
			termoMonitorado.Status = 0
		}

		newID, err := CreateTermoMonitorado(db, termoMonitorado)
		if err != nil {
			log.Println(err.Error())
			return
		}

		for _, variacao := range item.Variacoes.Item {
			var termoVariacao TermoVariacao
			termoVariacao.TermoID = newID
			termoVariacao.Variacao = variacao.Termo
			_, err := CreateTermoVariacao(db, termoVariacao)
			if err != nil {
				log.Println(err.Error())
				return
			}
		}
	}
}

type Nomes struct {
	Item []struct {
		CodNome        int64  `json:"codNome"`
		CodEscritorio  int64  `json:"codEscritorio"`
		Nome           string `json:"nome"`
		Oab            string `json:"oab"`
		DataCadastro   string `json:"dataCadastro"`
		TermosBloqueio string `json:"termosBloqueio"`
		Variacoes      struct {
			Item []struct {
				Termo          string `json:"Termo"`
				TermosBloqueio string `json:"TermosBloqueio"`
			} `json:"Item"`
		} `json:"variacoes"`
		Abrangencia struct {
			Item []string `json:"item"`
		} `json:"Abrangencia"`
	} `json:"Item"`
}

type TermoMonitorado struct {
	ID            int64  `json:"id"`
	EscID         int64  `json:"esc_id"`
	Termo         string `json:"termo"`
	CustomID      string `json:"custom_id"`
	SolucionareID int64  `json:"solucionare_id"`
	EscavadorID   int64  `json:"escavador_id"`
	DigestoID     int64  `json:"digesto_id"`
	Status        int    `json:"status"`
	UUID          string `json:"uuid"`
	CodAdv        int
	Variacoes     []*TermoVariacao    `json:"variacoes"`
	Diarios       []*DiarioMonitorado `json:"diarios"`
	CreatedAt     time.Time           `json:"created_at"`
	UpdatedAt     time.Time           `json:"updated_at"`
}

func CreateTermoMonitorado(db *sql.DB, termo TermoMonitorado) (int64, error) {
	query := "INSERT INTO projud_dados.termos_monitorados (id_esc, termo, id_custom, id_escavador, id_solucionare, id_digesto, status, uuid, cod_adv) VALUES(?,?,?,?,?,?,?,?,?);"
	stmt, err := db.Prepare(query)
	if err != nil {
		return -1, err
	}
	res, err := stmt.Exec(
		termo.EscID,
		termo.Termo,
		termo.CustomID,
		termo.EscavadorID,
		termo.SolucionareID,
		termo.DigestoID,
		termo.Status,
		termo.UUID,
		termo.CodAdv,
	)
	if err != nil {
		return -1, err
	}
	if err != nil {
		return -1, err
	}
	defer stmt.Close()
	return res.LastInsertId()
}

type TermoVariacao struct {
	ID       int64
	TermoID  int64
	Variacao string
}

func CreateTermoVariacao(db *sql.DB, termoVariacao TermoVariacao) (int64, error) {
	query := "INSERT INTO projud_dados.termos_monitorados_variacoes (id_termo, variacao) VALUES(?,?);"
	stmt, err := db.Prepare(query)
	if err != nil {
		return -1, err
	}
	defer stmt.Close()
	res, err := stmt.Exec(
		termoVariacao.TermoID,
		termoVariacao.Variacao,
	)
	if err != nil {
		return -1, err
	}
	if err != nil {
		return -1, err
	}
	defer stmt.Close()
	return res.LastInsertId()
}

type DiarioMonitorado struct {
	ID               int64
	EscID            int64
	TermoID          int64
	DiarioID         int64
	DiarioFornecedor string
	Status           int
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type Escritorio struct {
	CodEscritorio     int    `json:"codEscritorio"`
	Area              int    `json:"area"`
	Nome              string `json:"nome"`
	Senha             string `json:"senha"`
	Endereco          string `json:"endereco"`
	Cidade            string `json:"cidade"`
	Estado            string `json:"estado"`
	Cep               string `json:"cep"`
	PerfilContratante string `json:"perfilContratante"`
	Telefone          string `json:"telefone"`
	DataCadastro      string `json:"dataCadastro"`
	Bloqueado         bool   `json:"bloqueado"`
}
