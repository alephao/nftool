package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime"
	"sync"

	"github.com/spf13/cobra"
)

var (
	openseaCmd = &cobra.Command{
		Use:   "opensea",
		Short: "commands to interact with opensea",
	}

	// opensea update metadata
	openseaUpdateContractAddress string
	openseaUpdateFrom            int
	openseaUpdateTo              int
	openseaUpdateMaxCores        int
	openseaUpdateCmd             = &cobra.Command{
		Use:   "update",
		Short: "Asks opensea to update the metadata of your collection.",
		Example: `nftool opensea update \
	--contract 0x0000000000000000000000000000000000000000 \
	--from 0 \
	--to 1000 \
	--parallel 4`,
		Run: func(cmd *cobra.Command, args []string) {
			uptadeMetadataInRangeParallel(openseaUpdateContractAddress, openseaUpdateFrom, openseaUpdateTo, openseaUpdateMaxCores)
		},
	}
)

func init() {
	// update metadata on opensea
	openseaUpdateCmd.Flags().StringVar(&openseaUpdateContractAddress, "contract", "", "address of the contract in mainnet")
	openseaUpdateCmd.MarkFlagRequired("contract")
	openseaUpdateCmd.Flags().IntVar(&openseaUpdateFrom, "from", 0, "the starting id for a the range of ids you want to update")
	openseaUpdateCmd.MarkFlagRequired("from")
	openseaUpdateCmd.Flags().IntVar(&openseaUpdateTo, "to", 0, "the last id for a the range of ids you want to update")
	openseaUpdateCmd.MarkFlagRequired("to")
	openseaUpdateCmd.Flags().IntVar(&openseaUpdateMaxCores, "parallel", 1, "divide requests between your cpus")
	openseaCmd.AddCommand(openseaUpdateCmd)
}

func uptadeMetadataInRangeParallel(contractAddress string, from, to int, maxCores int) {
	numCpu := runtime.NumCPU()
	cores := min(maxCores, numCpu)

	numberOfTokensToUpdate := to - from
	chunkSize := numberOfTokensToUpdate / cores
	fmt.Printf("cores: %d\nchunk: %d\ntotal: %d\n", cores, chunkSize, numberOfTokensToUpdate)

	var wg sync.WaitGroup

	for i := 0; i < cores; i++ {
		wg.Add(1)

		coreFrom := (i * chunkSize) + from
		coreTo := coreFrom + chunkSize

		go func() {
			uptadeMetadataInRange(contractAddress, coreFrom, coreTo)
			wg.Done()
		}()
	}

	wg.Wait()
}

func uptadeMetadataInRange(contractAddress string, from, to int) {
	for i := from; i < to; i++ {
		err := updateMetadataOf(contractAddress, i)
		if err != nil {
			fmt.Printf("request to update #%d failed: %s\n", i, err.Error())
		} else {
			fmt.Printf("requested update for #%d\n", i)
		}
	}
}

func updateMetadataOf(contractAddress string, id int) error {
	url := fmt.Sprintf("https://api.opensea.io/api/v1/asset/%s/%d/?force_update=true", contractAddress, id)

	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("failed to update with contract '%s' and id '%d': %s", contractAddress, id, err.Error())
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("invalid response when trying to update with contract '%s' and id '%d': %s", contractAddress, id, err.Error())
	}

	b := map[string]interface{}{}
	err = json.Unmarshal(body, &b)
	if err != nil {
		return fmt.Errorf("failed to unmarshall response when trying to update with contract '%s' and id '%d': %s", contractAddress, id, err.Error())
	}

	if _, ok := b["token_id"]; !ok {
		return fmt.Errorf("invalid response from opensea for contract '%s' and id '%d' (the token might not exist)", contractAddress, id)
	}

	return nil
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
