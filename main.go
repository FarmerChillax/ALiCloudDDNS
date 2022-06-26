package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/FarmerChillax/ALiCloudDDNS/client"
)

func main() {
	var stop string
	go func() {
		log.Println("[Start] 程序运行中, 按任意键关闭...")
		fmt.Scanln(&stop)
		os.Exit(0)
	}()
	ddnsClient := client.New()
	log.Println("初始化 ddns 客户端成功:", *ddnsClient)
	for {
		ddnsClient.Run()
		time.Sleep(10 * time.Second)
	}
}
