package helpers

import (
	"html/template"
	"net/http"
)

func RenderizeHtml(dados interface{}, caminho string, w http.ResponseWriter) error {

  parse, err := template.ParseFiles(caminho)

  if err != nil {
          return err
  }

  templateHtml := template.Must(parse, err) //pasta do boleto html

  w.Header().Set("Content-Type", "text/html; charset=utf8") //precisa disso para renderizar o html

  if err := templateHtml.Execute(w, dados); err != nil { //funcao que renderiza o html passando os dados
          return err
  }
  return nil
}
