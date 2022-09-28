package client

import (
	"context"
	"log"

	"github.com/FarmerChillax/ALiCloudDDNS/agent"
	"github.com/FarmerChillax/ALiCloudDDNS/config"
	"github.com/FarmerChillax/ALiCloudDDNS/notice"
	"github.com/FarmerChillax/ALiCloudDDNS/proto/ddns_server"
	"github.com/google/uuid"
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
	HeartServerAddress string
	uuid               string
}

func New(config *config.DDNSConfig) *DDNSClient {
	var ddnsClient *DDNSClient
	// 初始化 DNS Agent
	// 当前版本只做阿里云
	aliAgent := agent.NewALiAgent(config)
	// 获取当前解析得 ip
	dnsRecordIp, err := aliAgent.GetRecordIp()
	if err != nil {
		log.Fatalf("获取阿里云记录失败, err: %v", err)
	}
	// 初始化 Notice
	notice := notice.New(config.NoticeUrl)
	// 初始客户端 UUID
	u := uuid.New()
	ddnsClient = &DDNSClient{
		Agent:     aliAgent,
		DnsHostIp: dnsRecordIp,
		Notice:    notice,
		uuid:      u.String(),
	}

	// 用于获取本机 IP 的节点
	if config.ServerAddr != "" {
		// 通过 gRPC 长连接的方式维护心跳
		ddnsClient.HeartServerAddress = config.ServerAddr
	} else {
		// 通过轮询的方式获取本机 IP 的节点
		ddnsClient.GetCurrentIpClient = NewGetIpClient()
	}

	return ddnsClient
}

func (d *DDNSClient) Run(ctx context.Context) (err error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("recover err:", err)
		}
	}()

	if d.HeartServerAddress != "" {
		// 通过 gRPC 双向流维护心跳
		d.HeartBeat(ctx)
	}

	if d.GetCurrentIpClient != nil {
		// 通过轮询维护心跳
		return d.LongPoll(ctx)
	}
	return nil
}

func (d *DDNSClient) HeartBeat(ctx context.Context) {
	cc, err := grpc.DialContext(ctx, d.HeartServerAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return
	}

	defer cc.Close()

	dsc := ddns_server.NewDdnsServerClient(cc)
	ds, err := dsc.HeartBeatServer(ctx)
	if err != nil {
		return
	}

	go func() {
		// heartBeat, err := ds.Recv()
		// if err != nil {
		// 	log.Println("grpc HeartBeat recv err:", err)
		// }
		//  = heartBeat.GetIp()
	}()
	for {
		err := ds.Send(&ddns_server.HeartBeatClient{Uuid: d.uuid})
		if err != nil {
			log.Println("grpc HeartBeat send err:", err)
		}
	}

}

func (d *DDNSClient) sendHeartBeat(ctx context.Context, msg chan *ddns_server.HeartBeatClient) {

}

func (d *DDNSClient) LongPoll(ctx context.Context) error {
	currentIp, err := d.GetCurrentIpClient.Get()
	if err != nil {
		log.Println(err.Error())
		return err
	}
	if d.DnsHostIp != currentIp {
		// DNS 解析的 IP 与本机不符
		// 则更新解析 IP
		ok, err := d.Agent.Update(currentIp)
		if err != nil {
			log.Printf("更新解析 IP 出错, err: %v\n", err)
			d.Notice.Push(d.DnsHostIp, currentIp, "error")
			return err
		}
		if ok {
			log.Printf("[SUCCESS] 更新解析成功, %s -> %s", d.DnsHostIp, currentIp)
			d.Notice.Push(d.DnsHostIp, currentIp, "success")
			d.DnsHostIp = currentIp
		}
	} else {
		log.Println("IP 未发生变更, 无需更改...")
	}
	return nil
}
