package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "可视化配置 ddns",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("---开始设置 config.json ---")
	},
}
