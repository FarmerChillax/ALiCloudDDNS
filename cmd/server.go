package cmd

import (
	"net"

	"github.com/FarmerChillax/ALiCloudDDNS/log"
	"github.com/spf13/cobra"
)

var address = "127.0.0.1:5000"

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "fddns server",
	Run: func(cmd *cobra.Command, args []string) {
		lis, err := net.Listen("tcp", address)
		if err != nil {
			log.FileLog
		}
	},
}
