## 阿里云 DNS 动态解析

家里的宽带开通了公网ip，想进行域名解析， 但是路由器有没有ddns功能，因此创建该项目， 利用阿里云注册的备案域名和域名解析API，进行动态域名解析

--------------------------

## 使用

启动时 指定`access id` 和 `access key`

```
./ddns --accessId xxx --accessKey xxxx --domain my.domain.com --refresh 30
```

参数说明：

```
accessId: aliyun access id #注意创建的ram用户需要给aliyun dns 访问权限
accessKey: aliyun access key
domain:  需要解析的域名
refresh: 刷新检查ip间隔 30s

```


### 参考

https://help.aliyun.com/knowledge_detail/39863.html?spm=a2c4g.11186623.6.624.12f91550pQYrFU



