package ialidns

import (
	"log"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/alidns"
)

var AccessKeyID = ""
var AccessKeySecret = ""

var client *alidns.Client

func init() {
	var err error
	client, err = alidns.NewClientWithAccessKey("cn-hangzhou", AccessKeyID, AccessKeySecret)
	if err != nil {
		log.Fatal(err)
	}
}

func GetAlidnsClient() *alidns.Client {
	return client
}
