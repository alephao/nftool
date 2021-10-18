package fs

import (
	"github.com/alephao/nftool/traits/domain"
)

func loadMultipleTraitCollectionFiles(paths []string) ([]domain.TraitCollection, error) {
	allCollections := []domain.TraitCollection{}

	for _, path := range paths {
		attrs, err := LoadTraitCollectionFromFile(path)
		if err != nil {
			return nil, err
		}
		allCollections = append(allCollections, attrs)
	}

	return allCollections, nil
}
