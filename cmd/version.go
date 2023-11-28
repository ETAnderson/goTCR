package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use: 	"version",
	Short:	"Print the version number of GoTCR",
	Long:	"All software has versions. This is GoTCR's",
	Run:	func(cmd *cobra.Command, args []string) {
		fmt.Println("Golang Test And Commit or Reversion v0.1 -- HEAD")
	},
}
