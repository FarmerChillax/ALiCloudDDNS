package client

import (
	"context"
	"fmt"
	"log"
	"time"

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
	heartBeatRetry     int8
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
	}
	// 通过轮询的方式获取本机 IP 的节点
	ddnsClient.GetCurrentIpClient = NewGetIpClient()

	return ddnsClient
}

func (d *DDNSClient) Run(ctx context.Context) (err error) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("recover err:", err)
		}
	}()

	fmt.Println("heartBeat server address:", d.HeartServerAddress, d.HeartServerAddress != "" && d.heartBeatRetry < 10)
	// 如果存在自定义 ddns 心跳服务器，则优先使用
	if d.HeartServerAddress != "" && d.heartBeatRetry < 10 {
		// 通过 gRPC 双向流维护心跳
		err = d.HeartBeat(ctx)
		if err != nil {
			log.Printf("grpc heartBeat err: %v", err)
			d.heartBeatRetry++
			return err
		}
	}

	if d.GetCurrentIpClient != nil {
		// 通过轮询维护心跳
		if d.HeartServerAddress != "" && d.heartBeatRetry == 10 {
			log.Printf("grpc heartBeat retry count max, using longPoll heartBeat.")
			log.Println("restart to retry grpc heartBeat.")
		}
		return d.LongPoll(ctx)
	}
	return nil
}

func (d *DDNSClient) HeartBeat(ctx context.Context) error {
	cc, err := grpc.DialContext(ctx, d.HeartServerAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("grpc dial err:", err)
		return err
	}

	defer cc.Close()

	dsc := ddns_server.NewDdnsServerClient(cc)
	ds, err := dsc.HeartBeatServer(ctx)
	if err != nil {
		log.Println("init heartBeat client err:", err)
		return err
	}
	heartBeatResp, err := ds.Recv()
	if err != nil {
		log.Println("recv heartBeat err:", err)
		return err
	}
	log.Println("recv:", heartBeatResp)
	if d.DnsHostIp != heartBeatResp.GetIp() && heartBeatResp.Ip != "" {
		ok, err := d.Agent.Update(heartBeatResp.GetIp())
		if err != nil {
			log.Printf("更新解析 IP 出错, err: %v\n", err)
			d.Notice.Push(d.DnsHostIp, heartBeatResp.Ip, "error")
			return err
		}
		if ok {
			log.Printf("[SUCCESS] 更新解析成功, %s -> %s", d.DnsHostIp, heartBeatResp.Ip)
			// 推送更新成功提升
			d.Notice.Push(d.DnsHostIp, heartBeatResp.Ip, "success")
			// 更新 dns 解析记录，重置重试次数
			d.DnsHostIp = heartBeatResp.Ip
			d.heartBeatRetry = 0
		}
	}

	timer := time.NewTimer(time.Second * 1)
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-timer.C:
			// 发送心跳包
			err = ds.Send(&ddns_server.HeartBeatClient{Uuid: d.uuid})
			if err != nil {
				log.Println("heartBeat faild, err:", err)
				return err
			}
		}
		log.Printf("send heartBeat to %s success.", d.HeartServerAddress)
		timer.Reset(time.Second)
	}
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
