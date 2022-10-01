package client

import (
	"context"
	"errors"
	"log"

	"github.com/FarmerChillax/ALiCloudDDNS/proto/ddns_server"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	GRPCError = errors.New("grpc heartBeat error")
)

type baseGrpcClient struct{}

type HeartBeatClient struct {
	ddns_server.DdnsServer_HeartBeatServerClient
	conn *grpc.ClientConn
}

func NewHeartBeatClient(ctx context.Context, addr string) (*HeartBeatClient, error) {
	cc, err := grpc.DialContext(ctx, addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	ddnsHeartBeatClient := &HeartBeatClient{conn: cc}
	heartBeatClient := ddns_server.NewDdnsServerClient(cc)
	ds, err := heartBeatClient.HeartBeatServer(ctx)
	if err != nil {
		return nil, err
	}
	ddnsHeartBeatClient.DdnsServer_HeartBeatServerClient = ds

	return ddnsHeartBeatClient, nil
}

func (dc *HeartBeatClient) Close() error {
	return dc.conn.Close()
}

// 从服务端获取本机 ip
func (dc *HeartBeatClient) HandleRecv(ctx context.Context, recvChan chan<- *ddns_server.HeartBeat, errChan chan<- error) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			heartBeat, err := dc.Recv()
			if err != nil {
				errChan <- err
			}
			log.Println("recv a heartBeat.")
			recvChan <- heartBeat
		}
	}
}

// 发送本机信息给服务端
func (dc *HeartBeatClient) HandleSend(ctx context.Context, sendChan chan *ddns_server.HeartBeatClient, errChan chan error) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			err := dc.Send(<-sendChan)
			if err != nil {
				errChan <- err
			}
			log.Println("send a heartBeat.")
		}
	}
}
