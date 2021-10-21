package cloud

import (
	"testing"
)

func TestAliyun(t *testing.T) {
	server := CreateAliyun(AliyunConfig{
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
