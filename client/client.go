package client

type DNSAgent interface {
	GetIp(string) string
	Update(string) (bool, error)
}

type DDNSClient struct {
	Agent         DNSAgent
	DnsHostIp     string
	CurrentHostIp string
}

func New() *DDNSClient {
	// 当前版本只做阿里云
	aliAgent := ALiDNSClient{}
	ddnsClient := DDNSClient{
		Agent: &aliAgent,
	}
	return &ddnsClient
}
