package main

import (
	"fmt"

	"github.com/FarmerChillax/ALiCloudDDNS/config"
)

func main() {
	fmt.Println(*config.DDNSConf)
	// var stop string
	// go func() {
	// 	log.Println("[Start] 程序运行中, 按任意键关闭...")
	// 	fmt.Scanln(&stop)
	// 	os.Exit(0)
	// }()
	// ddnsClient := client.New()
	// log.Println("初始化 ddns 客户端成功:", *ddnsClient)
	// for {
	// 	ddnsClient.Run()
	// 	time.Sleep(10 * time.Minute)
	// }
}
