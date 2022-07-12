package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var template = `
[Unit]
Description=fddns client

[Service]
Type=simple
WorkingDirectory=%s 
ExecStart= %s # ./fddns -c <your config filename>
KillMode=process
Restart=on-failure
RestartSec=3s

[Install]
WantedBy=multi-user.target
`

var serviceCmd = &cobra.Command{
	Use:   "service",
	Short: "系统守护进程配置",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(template)
		ex, err := os.Executable()
		if err != nil {
			log.Fatalf("export service file err: %v", err)
		}
		dir := filepath.Dir(ex)
		serviceFile := fmt.Sprintf(template, dir, ex)
		ioutil.WriteFile(exportPath, []byte(serviceFile), os.ModePerm)
	},
}

func init() {
	// serviceCmd.Flags().StringVarP(&exportPath)
}
