## 阿里云 DNS 动态解析

家里的宽带开通了公网ip，想进行域名解析，利用域名服务商提供的api，进行动态解析。 

支持：阿里云域名， name.com


--------------------------

## 编译

Linux ARM:

```
env GOOS=linux GOARCH=arm go build -o myddns -mod=vendor 
```

Linux ARM64: 

```
env GOOS=linux GOARCH=arm64 go build -o myddns -mod=vendor 
```

Linux AMD64: 

```
env GOOS=linux GOARCH=amd64 go build -o myddns -mod=vendor 
```


## 使用

部署在家庭局域网内任意服务器上【树莓派或者路由器】。

查看帮助:

```
./myddns -h
```

阿里云域名：

```
./myddns --cloud aliyun --accessId xxx --accessKey xxxx --domain my.domain.com --refresh 30
```

name.com域名：

```
./myddns --cloud name.com --user xxx --token xxxx --domain my.domain.com --refresh 30
```

参数说明：

```
cloud: 服务商 [可选]域名服务商，支持: aliyun name.com , 默认为aliyun

// aliyun 配置
accessId:  aliyun access id    #注意创建的ram用户需要给aliyun dns 访问权限 必须
accessKey: aliyun access key 

// name.com 配置
user: 用户名
token： token 见: https://www.name.com/zh-cn/account/settings/api


domain:  需要解析的域名 #必须
refresh: 刷新检查ip间隔 30s #可选
```

  可以使用`systemd`进行进程管理, 新建文件`/etc/systemd/system/myddns.service`
  
```
[Unit]
Description=MyDDNS

[Service]
Type=simple
User=root
WorkingDirectory=/opt/ddns # 注意修改为实际二进制可执行文件所在的目录
ExecStart=/opt/ddns/myddns_arm --accessId xxxx --accessKey xx --domain my.domain.com #  注意修改为实际二进制可执行文件的路径
RestartSec=2
Restart=always

[Install]
WantedBy=multi-user.target
```

`systemd`启动服务

```
sudo systemctl enable myddns
sudo systemctl start myddns

#查看下启动是否成功
journalctl -u myddns.service -f  #  结束日志查看 ctrl+c
```

### 参考

- aliyun: https://help.aliyun.com/knowledge_detail/39863.html?spm=a2c4g.11186623.6.624.12f91550pQYrFU
- name.com: https://www.name.com/zh-cn/api-docs/
