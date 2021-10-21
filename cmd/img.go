package cmd

import (
	"github.com/alephao/nftool/imgen"
	"github.com/spf13/cobra"
)

var (
	// img
	imgCmd = &cobra.Command{
		Use:   "img",
		Short: "Manipulate images",
	}

	// img gen
	imgGenWidth          int
	imgGenHeight         int
	imgGenStartingIndex  int
	imgGenCollectionPath string
	imgGenConfigPath     string
	imgGenOutDir         string
	imgGenParallel       bool
	imgGenSaveAsPng      bool
	imgGenOnlyMissing    bool
	imgGenCmd            = &cobra.Command{
		Use:   "gen",
		Short: "Generate images from a collection.json and the layers folder",
		Long: `Generate images from a collection.json and the layers folder

nftool img gen --width 800 --height 800 --collection ./collection.json --config ./config.yaml --out ./generated-imgs --parallel --png`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := imgen.GenerateImagesFromCollectionAttributesJson(imgGenWidth, imgGenHeight, imgGenStartingIndex, imgGenCollectionPath, imgGenConfigPath, imgGenOutDir, imgGenParallel, imgGenSaveAsPng, imgGenOnlyMissing); err != nil {
				return err
			}
			return nil
		},
	}
)

func init() {
	// img gen
	imgGenCmd.Flags().IntVar(&imgGenWidth, "width", 0, "the width of the generated image")
	imgGenCmd.MarkFlagRequired("width")
	imgGenCmd.Flags().IntVar(&imgGenHeight, "height", 0, "the height of the generated image")
	imgGenCmd.MarkFlagRequired("height")
	imgGenCmd.Flags().IntVar(&imgGenStartingIndex, "startingIndex", 0, "the initial index of the image names, e.g.: you have a collection of 100 and want the names to be from 1000 to 1100, use --starrtingIndex 1000")
	imgGenCmd.Flags().StringVar(&imgGenCollectionPath, "collection", "", "path to the collection.json file")
	imgGenCmd.MarkFlagRequired("collection")
	imgGenCmd.Flags().StringVar(&imgGenConfigPath, "config", "", "path to the configuration file")
	imgGenCmd.MarkFlagRequired("config")
	imgGenCmd.Flags().StringVar(&imgGenOutDir, "out", "", "path to the folder where all imgs will be generated")
	imgGenCmd.MarkFlagRequired("out")
	imgGenCmd.Flags().BoolVar(&imgGenParallel, "parallel", false, "generate images in parallel using all your cores")
	imgGenCmd.Flags().BoolVar(&imgGenSaveAsPng, "png", false, "generate png images instead of jpg images")
	imgGenCmd.Flags().BoolVar(&imgGenOnlyMissing, "onlyMissing", false, "generate only the images that are missing in the output dir")

	imgCmd.AddCommand(imgGenCmd)
}
