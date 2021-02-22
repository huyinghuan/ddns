package ialidns

import (
	"fmt"
	"strings"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"
)

type Domain struct {
	DomainName string
	RR         string
	IP         string
}

/**
* 获取域名介些列表
 */
func GetDomainRecords(domain string) ([]alidns.Record, error) {
	client := GetAlidnsClient()
	request := alidns.CreateDescribeDomainRecordsRequest()
	request.DomainName = domain
	request.PageSize = requests.NewInteger(500)
	request.Type = "A"
	response, err := client.DescribeDomainRecords(request)
	if err != nil {
		return nil, err
	}
	return response.DomainRecords.Record, nil
}

func AddDomainRecord(domain Domain) error {
	client := GetAlidnsClient()
	request := alidns.CreateAddDomainRecordRequest()
	request.DomainName = domain.DomainName
	request.Value = domain.IP
	request.Type = "A"
	request.RR = domain.RR
	_, err := client.AddDomainRecord(request)
	return err
}

func Parse(domain string) (Domain, error) {
	arr := strings.Split(domain, ".")
	if len(arr) < 2 {
		return Domain{}, fmt.Errorf("域名配置错误: %s", domain)
	}

	domainName := strings.Join(arr[len(arr)-2:], ".")
	rr := strings.Join(arr[0:len(arr)-2], ".")
	return Domain{
		DomainName: domainName,
		RR:         rr,
	}, nil
}

func UpdateDomainRecord(id string, domain Domain) error {
	client := GetAlidnsClient()
	request := alidns.CreateUpdateDomainRecordRequest()
	request.RecordId = id
	request.Value = domain.IP
	request.Type = "A"
	request.RR = domain.RR
	_, err := client.UpdateDomainRecord(request)
	return err
}

func AddOrUpdateDomain(domain Domain) error {
	records, err := GetDomainRecords(domain.DomainName)
	if err != nil {
		return err
	}
	for _, record := range records {
		if record.RR == domain.RR {
			// ip 没有辩护，不需要重新解析
			if record.Value == domain.IP {
				return nil
			}
			return UpdateDomainRecord(record.RecordId, domain)
		}
	}
	return AddDomainRecord(domain)
}
