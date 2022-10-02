package cmd

import (
	"fmt"

	"github.com/FarmerChillax/ALiCloudDDNS/config"
	"github.com/spf13/cobra"
)

var (
	fileName string
	fileType string
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "配置文件",
	Run: func(cmd *cobra.Command, args []string) {
		d := config.New()
		d.Export(fmt.Sprintf("%s.%s", fileName, fileType))
	},
}

func init() {
	configCmd.Flags().StringVarP(&fileName, "name", "n", "config", "配置文件名")
	configCmd.Flags().StringVarP(&fileType, "type", "t", "json", "输出配置类型的后缀名")
}
