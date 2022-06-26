package client

import (
	"log"

	"github.com/FarmerChillax/ALiCloudDDNS/agent"
)

type DNSAgent interface {
	GetRecordIp() (string, error)
	Update(string) (bool, error)
}

type DDNSClient struct {
	Agent              DNSAgent
	GetCurrentIpClient *GetIpClient
	DnsHostIp          string
}

func New() *DDNSClient {
	// 用于获取本机 IP 的节点
	getIpClient := NewGetIpClient()
	// 当前版本只做阿里云
	aliAgent := agent.NewALiAgent()
	dnsRecordIp, err := aliAgent.GetRecordIp()
	if err != nil {
		log.Fatalf("获取阿里云记录失败, err: %v", err)
	}

	ddnsClient := DDNSClient{
		Agent:              aliAgent,
		GetCurrentIpClient: getIpClient,
		DnsHostIp:          dnsRecordIp,
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
		// DNS 解析的 IP 与本机不符
		// 则更新解析 IP
		ok, err := d.Agent.Update(currentIp)
		if err != nil {
			log.Fatalf("[INFO] 更新解析 IP 出错, err: %v\n", err)
			return
		}
		if ok {
			log.Printf("[SUCCESS] 更新解析成功, %s -> %s", d.DnsHostIp, currentIp)
			d.DnsHostIp = currentIp
		} else {
			log.Println("[INFO] 更新解析 IP 失败")
		}
	} else {
		log.Println("[INFO] IP 未发生变更, 无需更改...")
	}
}
