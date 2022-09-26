package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"strings"
)

type AliyunConfig struct {
	AccessKeyID     string `json:"accessId"`
	AccessKeySecret string `json:"accessKey"`
}

type NameComConfig struct {
	Username string `json:"username"`
	Token    string `json:"token"`
	API      string `json:"api"`
}

type DNSPodConfig struct {
	ID         string `json:"id"`
	LoginToken string `json:"loginToken"`
	API        string `json:"api"`
}

type Cloud struct {
	Aliyun AliyunConfig  `json:"aliyun"`
	Name   NameComConfig `json:"name"`
	DNSPod DNSPodConfig  `json:"dnspod"`
}

type Config struct {
	Debug    bool                `json:"debug"`
	Refresh  int                 `json:"refresh"`
	Cloud    Cloud               `json:"cloud"`
	Domains  map[string][]string `json:"domains"`
	GetIPAPI string              `json:"getIpApi"`
}

func (c *Config) Verity() {
	for cloud, domains := range c.Domains {
		switch cloud {
		case "aliyun":
			if c.Cloud.Aliyun.AccessKeyID == "" || c.Cloud.Aliyun.AccessKeySecret == "" {
				log.Fatalln("关键参数不能为空: accessId, accessKey")
			}
		case "name":
			if c.Cloud.Name.Username == "" || c.Cloud.Name.Token == "" {
				log.Fatalln("关键参数不能为空: username, token")
			}
		case "dnspod":
		default:
			log.Fatalf("不支持该域名服务商:%s", cloud)
		}
		log.Printf("服务商: %s 监控ip变动间隔: %ds  目标域名:\n   -- %s \n", cloud, conf.Refresh, strings.Join(domains, "\n   -- "))
	}
}

var conf *Config

func Setup(confPath string) *Config {
	body, err := ioutil.ReadFile(confPath)
	if err != nil {
		log.Println(err)
		log.Fatalf("读取配置文件错误:%s", confPath)
	}
	err = json.Unmarshal(body, &conf)
	if err != nil {
		log.Println(err)
		log.Fatal("格式化配置文件错误")
	}
	return conf
}

func Get() *Config {
	return conf
}
