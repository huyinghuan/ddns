## 阿里云 DNS 动态解析

家里的宽带开通了公网ip，想进行域名解析， 但是路由器有没有ddns功能，因此创建该项目， 利用阿里云注册的备案域名和域名解析API，进行动态域名解析

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

启动时 指定`access id` 和 `access key`

```
./myddns --accessId xxx --accessKey xxxx --domain my.domain.com --refresh 30
```

参数说明：

```
accessId: aliyun access id #注意创建的ram用户需要给aliyun dns 访问权限 必须
accessKey: aliyun access key #必须
domain:  需要解析的域名 #必须
refresh: 刷新检查ip间隔 30s #可选

```


### 参考

https://help.aliyun.com/knowledge_detail/39863.html?spm=a2c4g.11186623.6.624.12f91550pQYrFU



