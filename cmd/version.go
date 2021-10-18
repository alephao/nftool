package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of nftool",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("nftool v0.1")
	},
}
