package cmd

import "github.com/spf13/cobra"

var exportCmd = &cobra.Command{
	Use:   "export",
	Short: "导出相关文件",
}

func init() {
	exportCmd.AddCommand(configCmd)
}
