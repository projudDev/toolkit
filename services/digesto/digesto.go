package digesto

import "encoding/json"
import "bytes"
import "net/http"
import "errors"
import "strconv"
import "io/ioutil"

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
	if response.Status != "200 OK" {
		var e Erro
		if err = json.NewDecoder(response.Body).Decode(&e); err != nil {
			err = errors.New("Error: Fail converting digesto error" + err.Error())
			return
		}
		err = errors.New("Message: " + e.Message + " Status: " + strconv.Itoa(e.Status))
		return
	}
	err = json.NewDecoder(response.Body).Decode(&userCompany)
	return
}

func (this *Digesto) DeleteUserCompany(id int64) error {
	idS := strconv.FormatInt(id, 10)
	req, err := http.NewRequest("DELETE", this.Host+"/user_company/"+idS, nil)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+this.ApiKey)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	query, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != 204 {
		return errors.New(string(query))
	}

	return nil
}

func (this *Digesto) RestoreUserCompany(id int64) error {
	idS := strconv.FormatInt(id, 10)
	req, err := http.NewRequest("POST", this.Host+"/user_company/archive/"+idS+"/restore", nil)
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", "Bearer "+this.ApiKey)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return errors.New("Não foi possivel restaurar a empresa")
	}
	return nil
}

func (this *Digesto) UpdateUserCompany(id int64, ent *UserCompany) (userCompany *UserCompany, err error) {
	body, err := json.Marshal(ent)
	if err != nil {
		return
	}

	idS := strconv.FormatInt(id, 10)

	req, err := http.NewRequest("PATCH", this.Host+"/user_company/"+idS, bytes.NewReader(body))
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
	if response.Status != "200 OK" {
		var e Erro
		if err = json.NewDecoder(response.Body).Decode(&e); err != nil {
			err = errors.New("Error: Fail converting digesto error" + err.Error())
			return
		}
		err = errors.New("Message: " + e.Message + " Status: " + strconv.Itoa(e.Status))
		return
	}
	err = json.NewDecoder(response.Body).Decode(&userCompany)
	return
}

func (this *Digesto) CreateMonitoredPerson(ent *MonitoredPerson) (monitoredPerson *MonitoredPerson, err error) {
	body, err := json.Marshal(ent)
	if err != nil {
		return
	}
	req, err := http.NewRequest("POST", this.Host+"/monitored_person", bytes.NewReader(body))
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
	if response.Status != "200 OK" {
		var e Erro
		if err = json.NewDecoder(response.Body).Decode(&e); err != nil {
			err = errors.New("Error: Fail converting digesto error" + err.Error())
			return
		}
		err = errors.New("Message: " + e.Message + " Status: " + strconv.Itoa(e.Status))
		return
	}
	err = json.NewDecoder(response.Body).Decode(&monitoredPerson)
	return
}

func (this *Digesto) AdminApiToken(digestoID string) (userCompany *UserCompany, err error) {
	req, err := http.NewRequest("GET", this.Host+"/user_company/"+digestoID+"/admin_api_token", nil)
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
	if response.Status != "200 OK" {
		var e Erro
		if err = json.NewDecoder(response.Body).Decode(&e); err != nil {
			err = errors.New("Error: Fail converting digesto error" + err.Error())
			return
		}
		err = errors.New("Message: " + e.Message + " Status: " + strconv.Itoa(e.Status))
		return
	}
	err = json.NewDecoder(response.Body).Decode(&userCompany)
	return
}
