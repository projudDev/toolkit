package trf5

import (
	//"bytes"
	"context"
	//"net/http"
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
	"github.com/narrator69/captchure"
)

func startPje() {
	urlBase := "https://pje.trf5.jus.br/pje/ConsultaPublica/listView.seam"

	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var src string
	var ok bool

	erro := chromedp.Run(ctx,
		chromedp.Navigate(urlBase),
		chromedp.Sleep(2*time.Second),
		chromedp.AttributeValue(`#consultaPublicaForm\:captcha\:captchaImg`, "src", &src, &ok, chromedp.ByQuery),
	)

	if erro != nil {
		log.Fatal(erro)
	}

	urlImagem := "https://pje.trf5.jus.br/" + src

	respostaImagem, e := http.Get(urlImagem)
	if e != nil {
		log.Fatal(e)
	}
	defer respostaImagem.Body.Close()

	imagem, err := os.Create("captcha.png")
	if err != nil {
		log.Fatal(err)
	}
	defer imagem.Close()

	_, err = io.Copy(imagem, respostaImagem.Body)
	if err != nil {
		log.Fatal(err)
	}

	c := captchure.Captchure{ClientKey: "342dc79dca0ec09fe4bb925b2b5d2baa"}

	base64image, erro := captchure.LocalFileToBase64("captcha.png")
	if erro != nil {
		log.Fatal(erro)
	}

	taskParameters := make(map[string]interface{})

	taskParameters["minLength"] = 5

	imagemTexto, err := c.SolveImage(base64image, taskParameters)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(imagemTexto)

	cookieJar, err := cookiejar.New(nil)
	if err != nil {
		log.Fatal(err)
	}

	param := url.Values{}

	param.Add("AJAXREQUEST", "_viewRoot")
	param.Add("consultaPublicaForm:Processo:jurisdicaoSecaoDecoration:jurisdicaoSecao", "org.jboss.seam.ui.NoSelectionConverter.noSelectionValue")
	param.Add("consultaPublicaForm:Processo:ProcessoDecoration:Processo", "numeroProcesso")
	param.Add("consultaPublicaForm:nomeParte:nomeParteDecoration:nomeParte", "")
	param.Add("consultaPublicaForm:nomeParteAdvogado:nomeParteAdvogadoDecoration:nomeParteAdvogadoDecoration:nomeParteAdvogado", "")
	param.Add("consultaPublicaForm:classeJudicial:idDecorateclasseJudicial:classeJudicial", "")
	param.Add("consultaPublicaForm:classeJudicial:idDecorateclasseJudicial:j_id167_selection", "")
	param.Add("consultaPublicaForm:numeroCPFCNPJ:numeroCPFCNPJRadioCPFCNPJ:numeroCPFCNPJCNPJ", "")
	param.Add("consultaPublicaForm:numeroOABParte:numeroOABParteDecoration:numeroOABParteEstadoCombo", "org.jboss.seam.ui.NoSelectionConverter.noSelectionValue")
	param.Add("consultaPublicaForm:numeroOABParte:numeroOABParteDecoration:numeroOABParte", "")
	param.Add("consultaPublicaForm:numeroOABParte:numeroOABParteDecoration:j_id218", "")
	param.Add("consultaPublicaForm:captcha:j_id228:verifyCaptcha", imagemTexto)
	param.Add("consultaPublicaForm", "consultaPublicaForm")
	param.Add("autoScroll", "")
	param.Add("javax.faces.ViewState", "j_id1")
	param.Add("consultaPublicaForm:pesq", "consultaPublicaForm:pesq")
	param.Add("AJAX:EVENTS_COUNT", "1")

	primeiraRequisicao, err := http.NewRequest("POST", urlBase, strings.NewReader(param.Encode()))
	if err != nil {
		log.Fatal(err)
	}

	primeiraRequisicao.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	HtppClient := &http.Client{Jar: cookieJar,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		}}

	primeiraResposta, err := HtppClient.Do(primeiraRequisicao)
	if err != nil {
		log.Fatal(err)
	}

	defer primeiraResposta.Body.Close()

	segundaRequisicao, err := http.NewRequest("GET", urlBase, nil)
	if err != nil {
		log.Fatal(err)
	}

	segundaResposta, err := HtppClient.Do(segundaRequisicao)
	if err != nil {
		log.Fatal(err)
	}

	defer segundaResposta.Body.Close()

	html, err := goquery.NewDocumentFromReader(segundaResposta.Body)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(html.Find(".tableArea").Text())

}
