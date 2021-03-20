package myip

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
)

func isIP(ip string) bool {
	return net.ParseIP(ip) != nil
}

// https://jsonip.com/
func GetMyIP() (string, error) {
	resp, err := http.Get("https://jsonip.com")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return "", err
	}
	if v, ok := result["ip"]; ok {
		if vv, ok := v.(string); ok {
			if isIP(vv) {
				return vv, nil
			}
		}
	}

	return "", fmt.Errorf("解析接口结果失败:%s", string(body))
}

func GetMyIP1() (string, error) {
	resp, err := http.Get("https://api.ipify.org?format=json")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return "", err
	}
	if v, ok := result["ip"]; ok {
		if vv, ok := v.(string); ok {
			if isIP(vv) {
				return vv, nil
			}
		}
	}

	return "", fmt.Errorf("解析接口结果失败:%s", string(body))
}

func GetMyIP2() (string, error) {
	resp, err := http.Get("https://ip.cn/api/index?type=0&ip=")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return "", err
	}
	if v, ok := result["ip"]; ok {
		if vv, ok := v.(string); ok {
			if isIP(vv) {
				return vv, nil
			}
		}
	}
	return "", fmt.Errorf("解析接口结果失败:%v", result["ip"])
}
