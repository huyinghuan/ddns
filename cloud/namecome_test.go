package cloud

import (
	"testing"
)

func TestNameCome(t *testing.T) {
	server := CreateNameCom(NameComConfig{
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
