package cloud

import (
	"log"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"
	"github.com/huyinghuan/ddns/config"
)

type InsideAliyun struct {
	Client *alidns.Client
}

/**
* 获取域名介些列表
 */
func (cloud *InsideAliyun) GetDomainRecords(domain string) ([]Domain, error) {

	request := alidns.CreateDescribeDomainRecordsRequest()
	request.DomainName = domain
	request.PageSize = requests.NewInteger(500)
	request.Type = "A"
	response, err := cloud.Client.DescribeDomainRecords(request)
	if err != nil {
		return nil, err
	}
	list := response.DomainRecords.Record
	var queue []Domain
	for _, record := range list {
		queue = append(queue, Domain{
			RecordId: record.RecordId,
			RR:       record.RR,
			IP:       record.Value,
		})
	}
	return queue, nil
}

func (cloud *InsideAliyun) AddDomainRecord(domain Domain) error {
	request := alidns.CreateAddDomainRecordRequest()
	request.DomainName = domain.DomainName
	request.Value = domain.IP
	request.Type = "A"
	request.RR = domain.RR
	_, err := cloud.Client.AddDomainRecord(request)
	return err
}

func (cloud *InsideAliyun) UpdateDomainRecord(id string, domain Domain) error {
	request := alidns.CreateUpdateDomainRecordRequest()
	request.RecordId = id
	request.Value = domain.IP
	request.Type = "A"
	request.RR = domain.RR
	_, err := cloud.Client.UpdateDomainRecord(request)
	return err
}

// AddOrUpdateDomain
// @return isChange error
func (cloud *InsideAliyun) AddOrUpdateDomain(domain Domain) (bool, error) {
	// if insideAliyun == nil {

	// }
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

func CreateAliyun(conf config.AliyunConfig) *InsideAliyun {
	client, err := alidns.NewClientWithAccessKey("cn-hangzhou", conf.AccessKeyID, conf.AccessKeySecret)
	if err != nil {
		log.Fatal(err)
	}
	return &InsideAliyun{
		Client: client,
	}
}
