package client

import (
	"log"
)

type DNSAgent interface {
	GetIp(...string) string
	Update(string) (bool, error)
}

type DDNSClient struct {
	Agent              DNSAgent
	GetCurrentIpClient *GetIpClient
	DnsHostIp          string
}

func New(rr string) *DDNSClient {
	// 当前版本只做阿里云
	aliAgent := ALiDNSClient{}
	getIpclient := NewGetIpClient()
	ddnsClient := DDNSClient{
		Agent:              &aliAgent,
		GetCurrentIpClient: getIpclient,
		DnsHostIp:          aliAgent.GetIp(rr),
	}
	return &ddnsClient
}

func (d *DDNSClient) Run() {
	currentIp, err := d.GetCurrentIpClient.Get()
	if err != nil {
		log.Println(err.Error())
		return
	}
	if d.DnsHostIp != currentIp {
		// fmt.Println(d.DnsHostIp, currentIp)
		// DNS 解析的 IP 与本机不符
		// 则更新解析 IP
		ok, err := d.Agent.Update(currentIp)
		if err != nil {
			log.Fatalf("[INFO] 更新解析 IP 出错, err: %v\n", err)
			return
		}
		if !ok {
			log.Println("[INFO] 更新解析 IP 失败")
		}
	}
}
