package cmd

import (
	"github.com/spf13/cobra"
)

var clientCmd = &cobra.Command{
	Use:   "client",
	Short: "fddns client",
	Run: func(cmd *cobra.Command, args []string) {

	},
}
