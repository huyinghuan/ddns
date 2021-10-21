package main

import (
	"flag"
	"log"
	"strings"
	"time"

	"github.com/huyinghuan/ddns/cloud"
	"github.com/huyinghuan/ddns/myip"
)

var Version = "v1.0.0"
var BuildTime = "Dev"

type LastIpRecord struct {
	IP         string
	RecordTime time.Time
}

var lastIPMap = make(map[string]LastIpRecord)

func main() {
	// 阿里云配置
	var accessId, accessKey string
	// name.com 配置: https://www.name.com/zh-cn/api-docs
	var username, token string
	//  域名
	var domainName string
	var cloudName string

	flag.StringVar(&cloudName, "cloud", "aliyun", "[可选]域名服务商，支持: aliyun name.com , 默认为aliyun")
	flag.StringVar(&accessId, "accessId", "", "阿里云access id")
	flag.StringVar(&accessKey, "accessKey", "", "阿里云access key")
	flag.StringVar(&username, "username", "", "name.com 用户名")
	flag.StringVar(&token, "token", "", "name.com token")
	flag.StringVar(&domainName, "domain", "", "目标域名, 多个域名用逗号隔开")

	var fresh int
	flag.IntVar(&fresh, "refresh", 30, "监控ip变动时间间隔【建议30s以上，30s】")
	flag.Parse()
	log.Println("Program Version  : ", Version)
	log.Println("Program BuildTime: ", BuildTime)
	log.Println("Program Author   : ", "ec.huyinghuan@gmail.com")
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	if fresh <= 1 {
		fresh = 30
	}
	if domainName == "" {
		log.Fatalln("关键参数不能为空: domain")
	}

	tmpList := strings.Split(domainName, ",")
	targetDomainList := []string{}
	for _, item := range tmpList {
		item = strings.TrimSpace(item)
		if item == "" {
			continue
		}
		if _, err := cloud.ParseDomain(item); err != nil {
			log.Fatalln("域名配置错误:" + item)
		}
		targetDomainList = append(targetDomainList, item)
	}

	conf := cloud.Config{
		Aliyun: cloud.AliyunConfig{
			AccessKeyID:     accessId,
			AccessKeySecret: accessKey,
		},
		NameCom: cloud.NameComConfig{
			Username: username,
			Token:    token,
		},
		Refresh: fresh,
	}

	var cloudServer cloud.CloudServer
	if cloudName == "" {
		cloudName = "aliyun"
	}
	if cloudName == "aliyun" {
		if conf.Aliyun.AccessKeyID == "" || conf.Aliyun.AccessKeySecret == "" {
			log.Fatalln("关键参数不能为空: accessId, accessKey")
		}
		cloudServer = cloud.CreateAliyun(conf.Aliyun)
	} else if cloudName == "name.com" {
		if conf.NameCom.Username == "" || conf.NameCom.Token == "" {
			log.Fatalln("关键参数不能为空: username, token")
		}
		cloudServer = cloud.CreateNameCom(conf.NameCom)
	} else if cloudName == "dnspod" {

	} else {
		log.Fatalln("不支持该域名服务商")
	}

	log.Printf("服务商: %s 监控ip变动间隔: %ds  目标域名:\n   -- %s \n", cloudName, conf.Refresh, strings.Join(targetDomainList, "\n   -- "))

	timer := time.NewTicker(time.Duration(fresh) * time.Second)
	for {
		ip := myip.GetMyIP()
		if ip == "" {
			log.Println("获取本机ip失败")
			<-timer.C
			continue
		}
		for _, item := range targetDomainList {
			//时间内10分钟，ip没有不用查阿里云接口
			if v, ok := lastIPMap[item]; ok && v.IP == ip && v.RecordTime.Add(10*time.Minute).After(time.Now()) {
				continue
			}
			domain, _ := cloud.ParseDomain(item)
			domain.IP = ip
			if hasChange, err := cloudServer.AddOrUpdateDomain(domain); err != nil {
				log.Println(err)
				log.Printf("域名: %s 更新失败 \n", item)
			} else if hasChange {
				lastIPMap[item] = LastIpRecord{
					IP:         ip,
					RecordTime: time.Now(),
				}
				log.Printf("解析成功: %s ==> %s \n", domain.OriginDomain, domain.IP)
			} else {
				lastIPMap[item] = LastIpRecord{
					IP:         ip,
					RecordTime: time.Now(),
				}
			}
		}
		<-timer.C
	}
}
