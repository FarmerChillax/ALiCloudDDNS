package cmd

import (
	"fmt"

	"github.com/FarmerChillax/ALiCloudDDNS/config"
	"github.com/spf13/cobra"
)

var filename string

var setCmd = &cobra.Command{
	Use:   "set",
	Short: "生成 ddns 配置文件",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args)
		config.DDNSConf.Save(filename)
	},
}

func init() {
	fmt.Println("设置前的 config:", config.DDNSConf)
	setCmd.Flags().StringVarP(&filename,
		"file", "f", "config.json", "保存的文件名")
	setCmd.Flags().StringVarP(&config.DDNSConf.AccessKey,
		"key", "k", "YOUR ACCESS KEY ID", "阿里云 AccessKey")
	setCmd.Flags().StringVarP(&config.DDNSConf.AccessKeySecret,
		"secret", "s", "YOUR ACCESS KEY SECRET", "阿里云 AccessKeySecret")
	// config.DDNSConf.AccessKey
	setCmd.Flags().StringVarP(&config.DDNSConf.DomainName,
		"name", "n", "example.com", "动态解析的域名")

	setCmd.Flags().StringVarP(&config.DDNSConf.RR,
		"rr", "r", "@", "没想好叫啥，反正就是阿里云的 RR")

	setCmd.Flags().StringVarP(&config.DDNSConf.Type,
		"type", "t", "A", "需要解析的类型（与阿里云相对应）")
	setCmd.Flags().StringVarP(&config.DDNSConf.RegionId,
		"region", "", "cn-hangzhou", "解析的 region")
}
