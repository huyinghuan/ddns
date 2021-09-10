package cloud

import (
	"testing"

	"github.com/huyinghuan/ddns/config"
)

func TestNameCome(t *testing.T) {
	server := CreateNameCom(config.NameComConfig{
		Username: "",
		Token:    "",
		API:      "https://api.name.com",
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
