package ialidns

import (
	"fmt"
	"log"
	"strings"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"
)

type Domain struct {
	DomainName   string
	RR           string
	IP           string
	OriginDomain string
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
	if err == nil {
		log.Println("增加解析记录:", domain.OriginDomain, domain.IP)
	}
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
		DomainName:   domainName,
		RR:           rr,
		OriginDomain: domain,
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
	if err == nil {
		log.Println("修改解析记录成功:", domain.OriginDomain, domain.IP)
	}
	return err
}

// AddOrUpdateDomain
// @return isChange error
func AddOrUpdateDomain(domain Domain) (bool, error) {
	records, err := GetDomainRecords(domain.DomainName)
	if err != nil {
		return false, err
	}
	for _, record := range records {
		if record.RR == domain.RR {
			// ip 没有变化，不需要重新解析
			if record.Value == domain.IP {
				return false, nil
			}
			return true, UpdateDomainRecord(record.RecordId, domain)
		}
	}
	return true, AddDomainRecord(domain)
}
