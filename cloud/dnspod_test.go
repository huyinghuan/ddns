package cloud

import (
	"log"
	"testing"
)

func TestDNSPod(t *testing.T) {

	cloud := CreateDNSPod(DNSPodConfig{
		LoginToken: "",
		ID:         "258498",
	})

	list, err := cloud.GetDomainRecords("opying.com")
	if err != nil {
		t.Fatal(err)
	}
	log.Println(list)
}
