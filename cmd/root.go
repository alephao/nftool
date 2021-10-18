package cmd

import (
	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "nftool",
		Short: "A suite of tools to for NFT generative art.",
	}
	rootCmd.AddCommand(traitsCmd)
	rootCmd.AddCommand(rarityCmd)
	rootCmd.AddCommand(imgCmd)
	rootCmd.AddCommand(provenanceCmd)
	rootCmd.AddCommand(versionCmd)
	return rootCmd
}

// Execute executes the root command.
func Execute() error {
	rootCmd := NewRootCmd()
	return rootCmd.Execute()
}
