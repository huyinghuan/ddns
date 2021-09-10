package cloud

import (
	"testing"

	"github.com/huyinghuan/ddns/config"
)

func TestAliyun(t *testing.T) {
	server := CreateAliyun(config.AliyunConfig{
		AccessKeyID:     "",
		AccessKeySecret: "",
	})

	hasChange, err := server.AddOrUpdateDomain(Domain{
		RR:         "",
		IP:         "",
		DomainName: "",
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(hasChange)
}
