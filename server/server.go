package server

import (
	"context"
	"log"

	"github.com/FarmerChillax/ALiCloudDDNS/proto/ddns_server"
	"google.golang.org/grpc/peer"
)

type DdnsServer struct {
}

func NewDdnsServer() *DdnsServer {
	return &DdnsServer{}
}

func (ds *DdnsServer) HeartBeatServer(stream ddns_server.DdnsServer_HeartBeatServerServer) error {
	for {
		in, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Printf("HeartBeat uuid: %s", in.GetUuid())
		peer, ok := peer.FromContext(context.Background())
		// if ok {
		// 	log.Println(peer, ok)
		// }
		log.Println(peer, ok)
	}
	return nil
}
