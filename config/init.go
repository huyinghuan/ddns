package config

type AliyunConfig struct {
	AccessKeyID     string
	AccessKeySecret string
}

type NameComConfig struct {
	Username string
	Token    string
	API      string
}

type Config struct {
	Refresh int
	Aliyun  AliyunConfig
	NameCom NameComConfig
}
