package fs

import (
	"fmt"
	"strconv"

	"github.com/alephao/nftool/metadata/domain"
	traits_fs "github.com/alephao/nftool/traits/fs"
	"github.com/alephao/nftool/utils"
)

func GenerateMetadata(collectionPath, configPath, outDir string, use1155Pattern bool) error {
	traitCollection, err := traits_fs.LoadTraitCollectionFromFile(collectionPath)
	if err != nil {
		return fmt.Errorf("failed to load collection file at '%s': %s", collectionPath, err.Error())
	}

	config, err := traits_fs.LoadConfigFromFile(configPath)
	if err != nil {
		return fmt.Errorf("failed to load config file at '%s': %s", configPath, err.Error())
	}

	metadataItems := domain.GenerateMetadata(traitCollection, config.Name, config.Description, config.Image, config.ExternalLink)

	for i, item := range metadataItems {
		var fileName string
		if use1155Pattern {
			fileName = fmt.Sprintf("%064x", i)
		} else {
			fileName = strconv.Itoa(i)
		}
		err = utils.WriteFileAsJson(item, fmt.Sprintf("%s/%s", outDir, fileName))
		if err != nil {
			return err
		}
	}

	return nil
}
