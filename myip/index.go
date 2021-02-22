package myip

import (
	"io/ioutil"
	"log"
	"net/http"
)

func GetMyIp() (string, error) {
	resp, err := http.Get("https://www.taobao.com/help/getip.php")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	log.Println(string(body))
	return "", err
}
