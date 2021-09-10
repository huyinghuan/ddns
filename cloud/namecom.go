package cloud

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"github.com/huyinghuan/ddns/config"
)

type NameComItem struct {
	ID     int64  `json:"id,omitempty"`
	Domain string `json:"domainName,omitempty"`
	Host   string `json:"host,omitempty"`
	Type   string `json:"type,omitempty"`
	IP     string `json:"answer,omitempty"`
	TTL    int    `json:"ttl,omitempty"`
}

type NameComItemList struct {
	Records []NameComItem `json:"records"`
}

type NameComError struct {
	Message string `json:"message"`
	Details string `json:"details"`
}

type InsideNameCom struct {
	API      string
	Username string
	Token    string
}

var client = http.Client{
	Timeout: 3 * time.Second,
}

func (cloud *InsideNameCom) GetDomainRecords(domain string) ([]Domain, error) {
	api := fmt.Sprintf("%s/v4/domains/%s/records", cloud.API, domain)
	request, _ := http.NewRequest("GET", api, nil)
	request.SetBasicAuth(cloud.Username, cloud.Token)
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, fmt.Errorf(string(body))
	}
	var items NameComItemList
	if err := json.Unmarshal(body, &items); err != nil {
		return nil, nil
	}
	queue := []Domain{}
	for _, record := range items.Records {
		queue = append(queue, Domain{
			RecordId:   strconv.FormatInt(record.ID, 10),
			RR:         record.Host,
			IP:         record.IP,
			DomainName: record.Domain,
		})
	}
	return queue, nil
}
func (cloud *InsideNameCom) AddDomainRecord(domain Domain) error {
	api := fmt.Sprintf("%s/v4/domains/%s/records", cloud.API, domain.DomainName)
	item := NameComItem{
		Host: domain.RR,
		Type: "A",
		IP:   domain.IP,
		TTL:  300,
	}
	body, _ := json.Marshal(item)
	request, _ := http.NewRequest("POST", api, bytes.NewReader(body))
	request.Header.Set("Content-Type", "application/json")
	request.SetBasicAuth(cloud.Username, cloud.Token)
	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}
	if response.StatusCode != 200 {
		return fmt.Errorf(string(responseBody))
	}
	return nil
}
func (cloud *InsideNameCom) UpdateDomainRecord(id string, domain Domain) error {
	api := fmt.Sprintf("%s/v4/domains/%s/records/%s", cloud.API, domain.DomainName, id)
	item := NameComItem{
		Host: domain.RR,
		Type: "A",
		IP:   domain.IP,
		TTL:  300,
	}
	body, _ := json.Marshal(item)
	request, _ := http.NewRequest("PUT", api, bytes.NewReader(body))
	request.Header.Set("Content-Type", "application/json")
	request.SetBasicAuth(cloud.Username, cloud.Token)
	response, err := client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()
	responseBody, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}
	if response.StatusCode != 200 {
		return fmt.Errorf(string(responseBody))
	}
	return nil
}

func (cloud *InsideNameCom) AddOrUpdateDomain(domain Domain) (bool, error) {
	records, err := cloud.GetDomainRecords(domain.DomainName)
	if err != nil {
		return false, err
	}
	for _, record := range records {
		if record.RR == domain.RR {
			// ip 没有变化，不需要重新解析
			if record.IP == domain.IP {
				return false, nil
			}
			return true, cloud.UpdateDomainRecord(record.RecordId, domain)
		}
	}
	return true, cloud.AddDomainRecord(domain)
}

func CreateNameCom(conf config.NameComConfig) *InsideNameCom {
	api := "https://api.name.com"
	if conf.API != "" {
		api = conf.API
	}
	return &InsideNameCom{
		API:      api,
		Username: conf.Username,
		Token:    conf.Token,
	}
}
