package ialidns

import (
	"log"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"
)

var AccessKeyID = ""
var AccessKeySecret = ""
var DomainName = ""

var client *alidns.Client

type Config struct {
	AccessKeyID     string
	AccessKeySecret string
	Domain          string
	Refresh         int
}

var config *Config

func init() {
	config = &Config{
		AccessKeyID:     AccessKeyID,
		AccessKeySecret: AccessKeySecret,
		Domain:          DomainName,
		Refresh:         30,
	}
}

func GetConfig() *Config {
	return config
}

func initClient() {
	var err error
	client, err = alidns.NewClientWithAccessKey("cn-hangzhou", config.AccessKeyID, config.AccessKeySecret)
	if err != nil {
		log.Fatal(err)
	}
}

func GetAlidnsClient() *alidns.Client {
	if client == nil {
		initClient()
	}
	return client
}
