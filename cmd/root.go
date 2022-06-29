package cmd

import (
	"log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "fddns",
	Short: "fddns 是一个 ddns 客户端",
	Long: `fddns 是一个轻量 ddns 客户端
目前仅支持阿里云服务，后续有需要会做更多的云服务商支持`,
	// 如果有相关的 action 要执行，请取消下面这行代码的注释
	// Run: func(cmd *cobra.Command, args []string) {
	// 	fmt.Println("启动主程序")
	// },
}

func init() {
	rootCmd.AddCommand(setCmd)
	rootCmd.AddCommand(startCmd)
	rootCmd.AddCommand(demoCmd)
}

// Execute 将所有子命令添加到root命令并适当设置标志。
// 这由 main.main() 调用。它只需要对 rootCmd 调用一次。
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
