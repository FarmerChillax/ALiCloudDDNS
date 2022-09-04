package client

import (
	"context"
	"log"

	"github.com/FarmerChillax/ALiCloudDDNS/agent"
	"github.com/FarmerChillax/ALiCloudDDNS/config"
	"github.com/FarmerChillax/ALiCloudDDNS/notice"
	"github.com/FarmerChillax/ALiCloudDDNS/proto/ddns_server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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
	heartClient        *grpc.ClientConn
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
	// 用于获取本机 IP 的节点
	if config.ServerAddr != "" {
		conn, err := grpc.Dial(config.ServerAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		ddnsClient.RegisterHeartBeatClientConn(conn)
	}
	return ddnsClient
}

func (d *DDNSClient) RegisterHeartBeatClientConn(conn *grpc.ClientConn) {
	d.heartClient = conn
}

func (d *DDNSClient) Run() {
	defer func() {
		if err := recover(); err != nil {
			log.Fatalln("recover err:", err)
		}
	}()
	if d.heartClient != nil {
		client := ddns_server.NewDdnsServerClient(d.heartClient)
		client.HeartBeatServer(context.Background())
	}
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
