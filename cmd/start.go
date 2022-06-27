package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "启动 ddns 服务",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("--- ddns 服务已启动 ---")
	},
}
