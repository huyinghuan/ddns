package myip

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
)

func isIP(ip string) bool {
	return net.ParseIP(ip).To4() != nil
}

// https://jsonip.com/
func getMyIP() (string, error) {
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

func getMyIP1() (string, error) {
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

func getMyIP2() (string, error) {
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

func getIpFromCustomServer(api string) (string, error) {
	if api == "" {
		return "", fmt.Errorf("ip api get server empty")
	}
	resp, err := http.Get(api)
	if err != nil {
		return "", err
	}
	if resp.StatusCode != 200 {
		return "", fmt.Errorf("自定义接口获取ip失败: %d", resp.StatusCode)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		return "", err
	}
	if isIP(string(body)) {
		return string(body), nil
	}
	return "", fmt.Errorf("ip格式错误:%s", string(body))
}
func GetMyIP(api string) string {

	ip := ""
	if ipAddr, err := getIpFromCustomServer(api); err == nil {
		ip = ipAddr
	} else if ipAddr, err := getMyIP(); err == nil {
		ip = ipAddr
	} else if ipAddr, err = getMyIP1(); err == nil {
		ip = ipAddr
	} else if ipAddr, err = getMyIP2(); err == nil {
		ip = ipAddr
	}
	return ip
}
