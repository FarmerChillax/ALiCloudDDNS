package cmd

import (
	"log"
	"time"

	"github.com/FarmerChillax/ALiCloudDDNS/client"
	"github.com/FarmerChillax/ALiCloudDDNS/config"
	flog "github.com/FarmerChillax/ALiCloudDDNS/log"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "fddns",
	Short: "fddns 是一个 ddns 客户端",
	Long: `fddns 是一个轻量 ddns 客户端
目前仅支持阿里云服务，后续有需要会做更多的云服务商支持`,
	Version: flog.VERSION,
	// 如果有相关的 action 要执行，请取消下面这行代码的注释
	Run: func(cmd *cobra.Command, args []string) {
		if logFileName != "" {
			flog.SetLoggerWithFile(logFileName)
		}

		ddnsConfig := config.Get()
		ddnsClient := client.New(ddnsConfig)
		log.Printf("初始化 ddns 客户端成功, 客户端代理为: %s, 当前域名解析为: %s",
			ddnsClient.Agent.GetName(), ddnsClient.DnsHostIp)
		durationT := time.Minute * time.Duration(duration)
		timer := time.NewTimer(durationT)
		for ; true; <-timer.C {
			ddnsClient.Run()
			timer.Reset(durationT)
		}
	},
}

var (
	duration    int64
	logFileName string
)

func init() {

	rootCmd.AddCommand(exportCmd)
	rootCmd.Flags().Int64VarP(&duration, "duration", "d", 10, "更新检测间隔")
	rootCmd.Flags().StringVarP(&logFileName, "log-filename", "l", "", "日志文件名")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
