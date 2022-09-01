package server

import (
	"fmt"
	"io"
	"log"
	"strings"
	"time"

	"github.com/FarmerChillax/ALiCloudDDNS/config"
	"github.com/FarmerChillax/ALiCloudDDNS/notice"
	"github.com/FarmerChillax/ALiCloudDDNS/proto/ddns_server"
	"google.golang.org/grpc/peer"
)

type DdnsServer struct {
	notice  *notice.Notice
	clients map[string]string
}

func NewDdnsServer(config *config.DDNSConfig) *DdnsServer {
	notice := notice.New(config.NoticeUrl)
	fmt.Println("Init: NewDdnsServer")
	return &DdnsServer{
		notice:  notice,
		clients: make(map[string]string),
	}
}

func (ds *DdnsServer) HeartBeatServer(stream ddns_server.DdnsServer_HeartBeatServerServer) error {
	// 发送客户端请求 IP
	peer, ok := peer.FromContext(stream.Context())
	if !ok {
		return fmt.Errorf("解析 peer context 出错")
	}
	addrs := strings.Split(peer.Addr.String(), ":")
	if len(addrs) == 1 {
		return fmt.Errorf("addr 解析出错")
	}
	stream.Send(&ddns_server.HeartBeat{Ip: addrs[0]})
	ds.notice.Online(addrs[0], time.Now().Format(time.UnixDate))
	var uuid string = "未知 UUID"
	// 维持心跳
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			log.Println("Closed! EOF")
			ds.notice.Offline(uuid, addrs[0], time.Now().Format(time.UnixDate))
			delete(ds.clients, uuid)
			log.Println("clients:", ds.clients)
			return nil
		}
		if err != nil {
			log.Println("Recv err:", err)
			return err
		}
		uuid = in.GetUuid()
		ds.clients[uuid] = addrs[0]
		log.Printf("HeartBeat uuid: %s, ip: %s", in.GetUuid(), addrs[0])
		log.Println("clients:", ds.clients)
	}
}
