package cloud

import (
	"fmt"
	"strings"
)

type Domain struct {
	RecordId     string
	DomainName   string // example.com
	RR           string // www
	IP           string
	OriginDomain string // www.example.com
}

func ParseDomain(domain string) (Domain, error) {
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

type CloudServer interface {
	AddOrUpdateDomain(domain Domain) (bool, error)
}
