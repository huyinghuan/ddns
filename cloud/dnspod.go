package cloud

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type DNSPodResponseStatus struct {
	Code string `json:"code"`
	Msg  string `json:"message"`
}

type DNSPodItem struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Value string `json:"value"`
}
type DNSPodResponseDomain struct {
	Name string `json:"name"`
}
type DNSPodResponse struct {
	Status  DNSPodResponseStatus `json:"status"`
	Records []DNSPodItem         `json:"records"`
	Domain  DNSPodResponseDomain `json:"domain"`
}

type InsideDNSPod struct {
	API   string
	Token string
	ID    string
}

func (cloud *InsideDNSPod) getPublicParams() url.Values {
	params := url.Values{}
	params.Set("login_token", cloud.Token)
	params.Set("format", "json")
	params.Set("error_on_empty", "no")
	return params
}

func (cloud *InsideDNSPod) GetDomainRecords(domain string) ([]Domain, error) {
	api := fmt.Sprintf("%s/Record.List", cloud.API)

	params := cloud.getPublicParams()
	params.Set("domain", domain)
	request, _ := http.NewRequest("POST", api, bytes.NewReader([]byte(params.Encode())))
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
	var dnsPodResponse DNSPodResponse
	if err := json.Unmarshal(body, &dnsPodResponse); err != nil {
		return nil, nil
	}
	if dnsPodResponse.Status.Code != "1" {
		return nil, fmt.Errorf(dnsPodResponse.Status.Msg)
	}
	queue := []Domain{}
	for _, record := range dnsPodResponse.Records {
		queue = append(queue, Domain{
			RecordId:   record.ID,
			RR:         record.Name,
			IP:         record.Value,
			DomainName: dnsPodResponse.Domain.Name,
		})
	}
	return queue, nil
}
func (cloud *InsideDNSPod) AddDomainRecord(domain Domain) error {
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
	//	request.SetBasicAuth(cloud.Username, cloud.Token)
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
func (cloud *InsideDNSPod) UpdateDomainRecord(id string, domain Domain) error {
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
	//	request.SetBasicAuth(cloud.Username, cloud.Token)
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

func (cloud *InsideDNSPod) AddOrUpdateDomain(domain Domain) (bool, error) {
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

func CreateDNSPod(conf DNSPodConfig) *InsideDNSPod {
	api := "https://dnsapi.cn"
	if conf.API != "" {
		api = conf.API
	}
	return &InsideDNSPod{
		API:   api,
		Token: conf.LoginToken,
		ID:    conf.ID,
	}
}
