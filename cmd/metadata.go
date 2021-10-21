package cmd

import (
	"github.com/spf13/cobra"

	metadata "github.com/alephao/nftool/metadata/fs"
)

var (
	// metadata
	metadataCollectionPath string
	metadataConfigPath     string
	metadataOutDir         string
	metadata1155Pattern    bool
	metadataCmd            = &cobra.Command{
		Use:   "metadata",
		Short: "generate the final metadata using a collection generated from the traits command",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := metadata.GenerateMetadata(metadataCollectionPath, metadataConfigPath, metadataOutDir, metadata1155Pattern); err != nil {
				return err
			}
			return nil
		},
	}
)

func init() {
	// metadata
	metadataCmd.Flags().StringVar(&metadataCollectionPath, "collection", "", "path to the collection json")
	metadataCmd.MarkFlagRequired("collection")
	metadataCmd.Flags().StringVar(&metadataConfigPath, "config", "", "path to the config json")
	metadataCmd.MarkFlagRequired("config")
	metadataCmd.Flags().StringVar(&metadataOutDir, "out", "", "path to the folder where the generated metadata files will be created")
	metadataCmd.MarkFlagRequired("out")
	metadataCmd.Flags().BoolVar(&metadata1155Pattern, "erc1155", false, "if the file names generated should follow the ERC-1155 pattern: hex value padded with zeros to size 64")
}
