package main

import (
	_ "net/http/pprof"
	"time"

	"github.com/FarmerChillax/ALiCloudDDNS/cmd"
)

const VERSION = "0.2.0"

var duration = 10 * time.Second

func main() {
	// defer profile.Start().Stop()
	// defer profile.Start(profile.MemProfile, profile.MemProfileRate(1)).Stop()

	cmd.Execute()

	go func() {
		// http.ListenAndServe(":233", nil)
	}()

	// var stop string
	// go func() {
	// 	log.Println("[Start] 程序运行中, 按任意键关闭...")
	// 	fmt.Scanln(&stop)
	// 	os.Exit(0)
	// }()
	// fmt.Println("load:", config.DDNSConf)

	// ddnsClient := client.New(config.DDNSConf)
	// log.Println("初始化 ddns 客户端成功:", *ddnsClient)
	// timer := time.NewTimer(duration)
	// for ; true; <-timer.C {
	// 	ddnsClient.Run()
	// 	timer.Reset(duration)
	// }
}

func SetLogger() {
	// log.SetPrefix()
	// log.SetFlags(0)
	// log.SetOutput()
}
