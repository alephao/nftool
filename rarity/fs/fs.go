package fs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/alephao/nftool/rarity/domain"
	traits "github.com/alephao/nftool/traits/domain"
	"github.com/alephao/nftool/utils"
)

// Trait Rarity Report
func MakeTraitRarityReportFromCollectionFile(path, out string) error {
	attrsFile, err := ioutil.ReadFile(path)
	if err != nil {
		return fmt.Errorf("failed to read file at path '%s': %w", path, err)
	}

	traits := []traits.TraitGroup{}
	err = json.Unmarshal(attrsFile, &traits)
	if err != nil {
		return fmt.Errorf("failed to unmarshal attrs file at path '%s': %w", path, err)
	}

	report := domain.MakeTraitRarityReport(traits)

	err = utils.WriteFileAsJson(report, out)
	if err != nil {
		return err
	}

	return nil
}

// Collection Rarity Report
func MakeCollectionRarityReportFromJson(path, out string) error {
	collection, err := loadCollecitionFile(path)
	if err != nil {
		return err
	}
	rarity := domain.MakeCollectionRarityReport(collection)
	return utils.WriteFileAsJson(rarity, out)
}

func loadCollecitionFile(path string) (traits.TraitCollection, error) {
	var collection traits.TraitCollection
	err := utils.LoadJsonFileIntoStruct(path, &collection)
	if err != nil {
		return traits.TraitCollection{}, err
	}
	return collection, nil
}
