package main

import (
	"flag"
	"log"
	"time"

	"github.com/huyinghuan/ddns/ialidns"
	"github.com/huyinghuan/ddns/myip"
)

var Version = "v1.0.0"
var BuildTime = "Dev"

func main() {
	var accessId, accessKey, domainName string

	flag.StringVar(&accessId, "accessId", "", "阿里云access id")
	flag.StringVar(&accessKey, "accessKey", "", "阿里云access key")
	flag.StringVar(&domainName, "domain", "", "目标域名")

	var fresh int
	flag.IntVar(&fresh, "refresh", 30, "监控ip变动时间间隔【建议30s以上，30s】")

	flag.Parse()

	log.Println("Program Version  : ", Version)
	log.Println("Program BuildTime: ", BuildTime)
	log.Println("Program Author   : ", "ec.huyinghuan@gmail.com")

	config := ialidns.GetConfig()

	if accessId != "" {
		config.AccessKeyID = accessId
	}
	if accessKey != "" {
		config.AccessKeySecret = accessKey
	}
	if domainName != "" {
		config.Domain = domainName
	}
	if fresh > 1 {
		config.Refresh = fresh
	}

	if config.AccessKeyID == "" || config.AccessKeySecret == "" || config.Domain == "" {
		log.Fatalln("关键参数不能为空: accessId, accessKey, domain")
	}

	log.Printf("监控ip变动间隔: %ds  目标域名: %s\n", config.Refresh, config.Domain)

	domain, err := ialidns.Parse(config.Domain)
	if err != nil {
		log.Fatalf("域名配置错误: %s", config.Domain)
	}

	timer := time.NewTimer(time.Duration(fresh) * time.Second)
	lastestIp := ""
	for {
		ip := ""
		if ipAddr, err := myip.GetMyIP(); err == nil {
			ip = ipAddr
		} else if ipAddr, err = myip.GetMyIP1(); err == nil {
			ip = ipAddr
		} else if ipAddr, err = myip.GetMyIP2(); err == nil {
			ip = ipAddr
		}
		if ip == "" {
			log.Println("获取本机ip失败")
		} else if ip != lastestIp {
			domain.IP = ip
			if err := ialidns.AddOrUpdateDomain(domain); err != nil {
				log.Println(err)
			} else {
				lastestIp = ip
			}
		} else {
			log.Println("地址未发生变化:", ip)
		}
		<-timer.C
	}
}
