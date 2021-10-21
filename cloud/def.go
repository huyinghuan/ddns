package cloud

type AliyunConfig struct {
	AccessKeyID     string
	AccessKeySecret string
}

type NameComConfig struct {
	Username string
	Token    string
	API      string
}

type DNSPodConfig struct {
	ID         string
	LoginToken string
	API        string
}

type Config struct {
	Refresh int
	Aliyun  AliyunConfig
	NameCom NameComConfig
	DNSPod  DNSPodConfig
}
