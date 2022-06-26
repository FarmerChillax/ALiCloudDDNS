package main

import (
	"log"

	"github.com/FarmerChillax/ALiCloudDDNS/client"
)

func main() {
	ddnsClient := client.New()
	log.Println("初始化 ddns 客户端成功:", *ddnsClient)
	ddnsClient.Run()
}
