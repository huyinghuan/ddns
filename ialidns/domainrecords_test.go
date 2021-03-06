package ialidns

import (
	"encoding/json"
	"testing"
)

func TestGetDomainRecord(t *testing.T) {
	records, err := GetDomainRecords("dpying.com")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(len(records))
	for _, d := range records {
		b, _ := json.Marshal(d)
		t.Log(string(b))
	}
}

func TestAddOrUpdateDomain(t *testing.T) {
}
