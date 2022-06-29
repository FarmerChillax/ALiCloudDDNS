package cmd

import "github.com/spf13/cobra"

var serverCmd = &cobra.Command{
	Use:   "serve",
	Short: "start a ddns get ip server",
	Run:   func(cmd *cobra.Command, args []string) {},
}
