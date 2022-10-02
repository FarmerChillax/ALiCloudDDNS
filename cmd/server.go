package cmd

import (
	"log"
	"net"

	"github.com/FarmerChillax/ALiCloudDDNS/config"
	"github.com/FarmerChillax/ALiCloudDDNS/proto/ddns_server"
	"github.com/FarmerChillax/ALiCloudDDNS/server"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var address = "127.0.0.1:5000"

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "fddns server",
	Run: func(cmd *cobra.Command, args []string) {
		lis, err := net.Listen("tcp", address)
		if err != nil {
			log.Fatalf("net.Listen err: %v", err)
		}

		grpcServer := grpc.NewServer()

		ddnsConfig := config.Get()
		ddnsServer := server.NewDdnsServer(ddnsConfig)
		ddns_server.RegisterDdnsServerServer(grpcServer, ddnsServer)

		log.Printf("starting gRPC listener on port: %v", address)
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	},
}
