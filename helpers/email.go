package helpers

import (
	"bytes"
	"crypto/tls"
	"database/sql"
	"fmt"
	"html/template"
	"io/ioutil"
	"strings"

	"github.com/leekchan/accounting"
	"gopkg.in/gomail.v2"
)

type MailCC struct {
	Address string
	Name    string
}

type Mail struct {
	Sender  string
	To      []string
	Subject string
	Body    string
	CC      []MailCC
	Attach  string
}

type SmtpServer struct {
	Host     string
	Port     int
	User     string
	Password string
	TLS      bool
}

func NewSmtpServer(Conn *sql.DB, codSistema int64) (SmtpServer, error) {
	var Server SmtpServer
	type LocalServer struct {
		Host     *string
		Port     *int
		User     *string
		Password *string
		TLS      *bool
	}
	localServer := LocalServer{}
	query := "SELECT email_host, email_user, email_senha, email_tls, email_porta FROM projud_teste.sistemas WHERE Cod=?;"
	err := Conn.QueryRow(query, codSistema).Scan(&localServer.Host, &localServer.User, &localServer.Password, &localServer.Port, &localServer.TLS)
	if localServer.Host != nil {
		Server.Host = *localServer.Host
	}
	if localServer.Port != nil {
		Server.Port = *localServer.Port
	}
	if localServer.User != nil {
		Server.User = *localServer.User
	}
	if localServer.Password != nil {
		Server.Password = *localServer.Password
	}
	if localServer.TLS != nil {
		Server.TLS = *localServer.TLS
	}
	return Server, err
}

func SendMail(Mail Mail, SmtpServer SmtpServer) error {
	m := gomail.NewMessage()
	m.SetHeader("From", Mail.Sender)
	m.SetHeader("To", Mail.To...)
	if len(strings.Trim(Mail.Attach, "")) > 0 {
		m.Attach(Mail.Attach)
	}
	for _, cc := range Mail.CC {
		m.SetAddressHeader("Cc", cc.Address, cc.Name)
	}

	m.SetHeader("Subject", Mail.Subject)
	m.SetBody("text/html", Mail.Body)

	d := gomail.NewDialer(SmtpServer.Host, SmtpServer.Port, SmtpServer.User, SmtpServer.Password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: SmtpServer.TLS}

	return d.DialAndSend(m)

}

func (m *Mail) ParseTemplate(fileName string, data interface{}) error {
	ac := accounting.Accounting{Symbol: "R$ ", Precision: 2, Thousand: ".", Decimal: ","}

	funcMap := template.FuncMap{
		"FormatNumber": func(value *float64) string {
			if value == nil {
				return "—"
			}
			if *value == 0 {
				return "—"
			}
			if *value == 0.00 {
				return "—"
			}
			if *value == 0.0 {
				return "—"
			}
			s := fmt.Sprintf("%2.f", *value)
			myText := strings.Replace(s, ".", "", -1)
			if myText == "0" {
				return "—"
			}
			return myText
		},
		"FormatCurrency": func(value *float64) string {
			if value == nil {
				return "R$ —"
			}
			if *value == 0 {
				return "R$ —"
			}
			if *value == 0.00 {
				return "R$ —"
			}
			if *value == 0.0 {
				return "R$ —"
			}
			return ac.FormatMoney(*value)
		},
		"SafeHTML": func(value string) template.HTML {
			return template.HTML(value)
		},
		"FormataDataPublicacao": func(value string) string {
			return Date2_html(Str2Date(value))
		},
	}

	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}

	buffer := new(bytes.Buffer)
	tmpl, err := template.New("email").Funcs(funcMap).Parse(string(file))
	if err != nil {
		return err
	}
	if err := tmpl.Execute(buffer, data); err != nil {
		return err
	}

	m.Body = buffer.String()
	return nil
}

func (m *Mail) SetTemplateAttach(fileName string, data interface{}) error {
	ac := accounting.Accounting{Symbol: "R$ ", Precision: 2, Thousand: ".", Decimal: ","}

	funcMap := template.FuncMap{
		"FormatNumber": func(value *float64) string {
			if value == nil {
				return "—"
			}
			if *value == 0 {
				return "—"
			}
			if *value == 0.00 {
				return "—"
			}
			if *value == 0.0 {
				return "—"
			}
			s := fmt.Sprintf("%2.f", *value)
			myText := strings.Replace(s, ".", "", -1)
			if myText == "0" {
				return "—"
			}
			return myText
		},
		"FormatCurrency": func(value *float64) string {
			if value == nil {
				return "R$ —"
			}
			if *value == 0 {
				return "R$ —"
			}
			if *value == 0.00 {
				return "R$ —"
			}
			if *value == 0.0 {
				return "R$ —"
			}
			return ac.FormatMoney(*value)
		},
		"SafeHTML": func(value string) template.HTML {
			return template.HTML(value)
		},
		"FormataDataPublicacao": func(value string) string {
			return Date2_html(Str2Date(value))
		},
	}

	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		return err
	}

	buffer := new(bytes.Buffer)
	tmpl, err := template.New("email").Funcs(funcMap).Parse(string(file))
	if err != nil {
		return err
	}
	if err := tmpl.Execute(buffer, data); err != nil {
		return err
	}
	err = ioutil.WriteFile("publicacoes.html", buffer.Bytes(), 0)
	if err != nil {
		return err
	}
	m.Attach = "publicacoes.html"
	return nil
}
