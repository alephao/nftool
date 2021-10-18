package cmd

import (
	"github.com/spf13/cobra"

	traits "github.com/alephao/nftool/traits/fs"
)

var (
	// shared vars
	traitsPath  string
	traitsPaths []string
	traitsOut   string

	// traits
	traitsCmd = &cobra.Command{
		Use:   "traits",
		Short: "Manipulate traits metadata",
	}

	// traits dump
	traitsDumpCmd = &cobra.Command{
		Use:   "dump",
		Short: "Generate a yaml configuration file from a directory containing all the layers and following the layer naming guidelines.",
		Long: `Generate a yaml configuration file from a directory containing all the layers and following the layer naming guidelines.

Example: nftool traits dump --path /path/to/layers --out /path/to/config.yaml

Layer Naming Guidelines

1. Layer Names

Inside the folder, each directory is a Layer, and you should name them following the pattern:

00-Background
01-Body
02-Face
03-Facial Hair

"Body" will be on top of "Background"
"Face" on top of "Body"
"Facial Hair" on top of "Face"

2. Layer Variant Names and Rarity

Each variant should follow the pattern <rarity weight>-<variant name>.png

00-Background
├── 1-Gray.png
└── 5-Haunted Mansion.png
01-Body
├── 1-Cool Hoodie.png
└── 1-Shinning Armour.png

The rarity of the traits will be calculated as a weighted average. For the example above, here is how we calculate the chance of a variant showing up:

First sum all rarity weights: 1 (Gray) + 5 (Haunted Mansion) = 6
Divide the individal weight by the sum. 
	For Gray: 1/6 = 16.7%
	For Haunted Mansion: 5/6 = 83.3%

In other words, for every 6 images generated, 1 will have the Gray background and 5 will have the Haunted Mansion background.

3. Optional Layers

If you want a layer to be optional, add a "_" suffix to the folder like so:
	
	00-Background_

Then after running 'nftool traits dump', you can change the chance of a trait not beign selected in the generated yaml file, under the "optional_weight" property.`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := traits.GenerateTraitGroupDescription(traitsPath, traitsOut); err != nil {
				return err
			}
			return nil
		},
	}

	// traits make
	traitsMakeAmount int
	traitsMakeCmd    = &cobra.Command{
		Use:   "make",
		Short: "Generate a collection from a config file or a folder with the structured layers.",
		Long: `Generate a collection from a config file or a folder with the structured layers.

Example: nftool traits make --path path/to/layers --out ./collection.json --amount 8888`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := traits.GenerateTraitCollection(traitsPath, traitsOut, traitsMakeAmount); err != nil {
				return err
			}
			return nil
		},
	}

	// traits shuffle
	traitsShuffleCmd = &cobra.Command{
		Use:   "shuffle",
		Short: "shuffle a collection json",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := traits.ShuffleCollection(traitsPath); err != nil {
				return err
			}
			return nil
		},
	}

	// traits merge
	traitsMergeCmd = &cobra.Command{
		Use:   "merge",
		Short: "Merge multiple attributes files in one, ignoring collisions",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := traits.MergeCollections(traitsPaths, traitsOut); err != nil {
				return err
			}
			return nil
		},
	}

	// traits collisions
	traitsCollisionsCmd = &cobra.Command{
		Use:   "collisions",
		Short: "Find collisions from multiple collection files and generate a report",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := traits.FindCollisions(traitsPaths, traitsOut); err != nil {
				return err
			}
			return nil
		},
	}
)

func init() {
	// traits dump
	traitsDumpCmd.Flags().StringVar(&traitsPath, "layers", "", "path to where all the layer folders are")
	traitsDumpCmd.MarkFlagRequired("layers")
	traitsDumpCmd.Flags().StringVar(&traitsOut, "out", "", "output path to the config, should have a .yaml extension")
	traitsDumpCmd.MarkFlagRequired("out")
	traitsCmd.AddCommand(traitsDumpCmd)

	// traits make
	traitsMakeCmd.Flags().StringVar(&traitsPath, "config", "", "path the configuration yaml file")
	traitsMakeCmd.MarkFlagRequired("config")
	traitsMakeCmd.Flags().StringVar(&traitsOut, "out", "", "output path to the collection, should have a .json extension")
	traitsMakeCmd.MarkFlagRequired("out")
	traitsMakeCmd.Flags().IntVar(&traitsMakeAmount, "amount", 0, "the amount of items you want to generate")
	traitsMakeCmd.MarkFlagRequired("amount")
	traitsCmd.AddCommand(traitsMakeCmd)

	// traits shuffle
	traitsShuffleCmd.Flags().StringVar(&traitsPath, "path", "", "path to the collection json")
	traitsShuffleCmd.MarkFlagRequired("path")
	traitsCmd.AddCommand(traitsShuffleCmd)

	// traits merge
	traitsMergeCmd.Flags().StringArrayVar(&traitsPaths, "file", nil, "list of collection json files to be merged")
	traitsMergeCmd.MarkFlagRequired("file")
	traitsMergeCmd.Flags().StringVar(&traitsOut, "out", "", "path to save the new collection json")
	traitsMergeCmd.MarkFlagRequired("out")
	traitsCmd.AddCommand(traitsMergeCmd)

	// traits collisions
	traitsCollisionsCmd.Flags().StringArrayVar(&traitsPaths, "file", nil, "list paths to collections")
	traitsCollisionsCmd.MarkFlagRequired("file")
	traitsCollisionsCmd.Flags().StringVar(&traitsOut, "out", "", "path to save the collision report")
	traitsCollisionsCmd.MarkFlagRequired("out")
	traitsCmd.AddCommand(traitsCollisionsCmd)
}
