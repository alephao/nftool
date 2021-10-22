package cmd

import (
	rarity "github.com/alephao/nftool/rarity/fs"
	"github.com/spf13/cobra"
)

var (
	// Used for flags.
	rarityPath string
	rarityOut  string
	rarityCmd  = &cobra.Command{
		Use:   "rarity",
		Short: "Get rarity info about a collection",
	}

	// rarity traits
	rarityTraitsCmd = &cobra.Command{
		Use:   "traits",
		Short: "Generate traits rarity report",
		Long:  "The report contains the number of appearances of each trait in the collection.",
		Example: `nftool rarity traits \
	--collection ./collection.json \
	--out ./traits_rarity.json`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := rarity.MakeTraitRarityReportFromCollectionFile(rarityPath, rarityOut); err != nil {
				return err
			}
			return nil
		},
	}

	// rarity collection
	rarityCollectionCmd = &cobra.Command{
		Use:   "collection",
		Short: "Generate collection rarity report",
		Long:  "The report contains all the items in the collection, ordered by most rare from least rare.",
		Example: `nftool rarity collection \
	--collection ./out/collection.json \
	--out ./out/collection_rarity.json`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := rarity.MakeCollectionRarityReportFromJson(rarityPath, rarityOut); err != nil {
				return err
			}
			return nil
		},
	}
)

func init() {
	// rarity traits
	rarityTraitsCmd.Flags().StringVar(&rarityPath, "collection", "", "path to collection json")
	rarityTraitsCmd.MarkFlagRequired("collection")
	rarityTraitsCmd.Flags().StringVar(&rarityOut, "out", "", "output path for the trait report, should have a .json extension")
	rarityTraitsCmd.MarkFlagRequired("out")
	rarityCmd.AddCommand(rarityTraitsCmd)

	// rarity collection
	rarityCollectionCmd.Flags().StringVar(&rarityPath, "collection", "", "path to collection json")
	rarityCollectionCmd.MarkFlagRequired("collection")
	rarityCollectionCmd.Flags().StringVar(&rarityOut, "out", "", "output path for the collection report, should have a .json extension")
	rarityCollectionCmd.MarkFlagRequired("out")
	rarityCmd.AddCommand(rarityCollectionCmd)
}
