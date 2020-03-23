package digesto

import "encoding/json"
import "bytes"
import "net/http"

const (
	host = "https://op.digesto.com.br/api/"
)

type Digesto struct {
	ApiKey string
	Host   string
}

func NewClient(ApiKey string) *Digesto {
	return &Digesto{
		ApiKey: ApiKey,
		Host:   host,
	}
}

func (this *Digesto) CreateUserCompany(ent *UserCompany) (userCompany *UserCompany, err error) {
	body, err := json.Marshal(ent)
	if err != nil {
		return
	}
	req, err := http.NewRequest("POST", this.Host+"/user_company", bytes.NewReader(body))
	if err != nil {
		return
	}
	req.Header.Set("Authorization", "Bearer "+this.ApiKey)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		return
	}
	err = json.NewDecoder(response.Body).Decode(&userCompany)
	return
}

func (this *Digesto) CreateMonitoredPerson(ent *MonitoredPerson) (*MonitoredPerson, error) {
	return nil, nil
}
