package cmd

import (
	provenance "github.com/alephao/nftool/provenance/fs"
	"github.com/spf13/cobra"
)

var (
	provenanceImgsDir       string
	provenanceOut           string
	provenanceStartingIndex int
	provenanceCmd           = &cobra.Command{
		Use:   "provenance",
		Short: "Generate the provenace report of a collection",
		Long: `How provenance is generated:

First we generate the hash for each image using a SHA256 algorithm

Then we append each hash following a specific order. The order is the original order offseted by the "startingIndex". We calculate each new index of images with the following formula:

(tokenId + startingIndex) \% collectionSize

Usually the startIndex is a number that is randomly generated on-chain.

After appending all hashes, we hash the result using SHA-256 again and that's the "final proof". This proof is usually stored on-chain.

The provenance report contains
- final proof hash
- concatenated hashes
- hashes
- starting index`,
		Example: `nftool provenance \
	--imgs ./imgs \
	--startingIndex 123 \
	--out ./provenance.json`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := provenance.ProvenanceReportFromDir(provenanceImgsDir, provenanceOut, provenanceStartingIndex); err != nil {
				return err
			}
			return nil
		},
	}
)

func init() {
	provenanceCmd.Flags().StringVar(&provenanceImgsDir, "images", "", "path to the directory containing the generated images")
	provenanceCmd.MarkFlagRequired("images")
	provenanceCmd.Flags().StringVar(&provenanceOut, "out", "", "where to save the provenance report")
	provenanceCmd.MarkFlagRequired("out")
	provenanceCmd.Flags().IntVar(&provenanceStartingIndex, "startingIndex", 0, "startingIndex")
}
