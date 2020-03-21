package trf5

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

func startTrf5() {
	baseUrl := "http://www4.trf5.jus.br/cp/cp.do"
	response, err := http.PostForm(baseUrl, url.Values{
		"tipo":          {"xmlproc"},
		"ordenacao cpf": {"D"},
		"tipoproc":      {"T"},
		"filtro":        {"0298966-76.2016.4.05.0000"},
	})

	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", string(body))
}
