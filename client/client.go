package client

import (
	"log"

	"github.com/FarmerChillax/ALiCloudDDNS/agent"
	"github.com/FarmerChillax/ALiCloudDDNS/config"
	"github.com/FarmerChillax/ALiCloudDDNS/notice"
)

type DNSAgent interface {
	GetRecordIp() (string, error)
	Update(string) (bool, error)
	SetName(string)
	GetName() string
}

type DDNSClient struct {
	Agent              DNSAgent
	GetCurrentIpClient *GetIpClient
	DnsHostIp          string
	Notice             *notice.Notice
}

func New(config *config.DDNSConfig) *DDNSClient {
	var ddnsClient *DDNSClient
	// 用于获取本机 IP 的节点
	getIpClient := NewGetIpClient()
	// 初始化 DNS Agent
	// 当前版本只做阿里云
	aliAgent := agent.NewALiAgent(config)
	dnsRecordIp, err := aliAgent.GetRecordIp()
	if err != nil {
		log.Fatalf("获取阿里云记录失败, err: %v", err)
	}

	// 初始化 Notice
	notice := notice.New(config.NoticeUrl)
	ddnsClient = &DDNSClient{
		Agent:              aliAgent,
		GetCurrentIpClient: getIpClient,
		DnsHostIp:          dnsRecordIp,
		Notice:             notice,
	}
	return ddnsClient
}

func (d *DDNSClient) Run() {
	defer func() {
		if err := recover(); err != nil {
			log.Fatalln("recover err:", err)
		}
	}()
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
			log.Printf("更新解析 IP 出错, err: %v\n", err)
			d.Notice.Push(d.DnsHostIp, currentIp, "error")
			return
		}
		if ok {
			log.Printf("[SUCCESS] 更新解析成功, %s -> %s", d.DnsHostIp, currentIp)
			d.Notice.Push(d.DnsHostIp, currentIp, "success")
			d.DnsHostIp = currentIp
		}
	} else {
		log.Println("IP 未发生变更, 无需更改...")
	}
}
