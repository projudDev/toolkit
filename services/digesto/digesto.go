package digesto

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
	req.Header.Set("Authorization", "Bearer 708c3227-156a-4dba-87e4-4a0061a671d2")
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, err := http.client.Do(req)
	if err != nil {
		return
	}
	err = json.NewDecoder(response.Body).Decode(&userCompany)
	return
}

func (this *Digesto) CreateMonitoredPerson(ent *MonitoredPerson) (*MonitoredPerson, error) {
	return nil, nil
}
