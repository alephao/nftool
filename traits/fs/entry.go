package fs

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/alephao/nftool/traits/domain"
	"github.com/alephao/nftool/utils"
)

type Config struct {
	Name         string                       `json:"name" yaml:"name"`
	Description  string                       `json:"description" yaml:"description"`
	ExternalLink string                       `json:"external_link" yaml:"external_link"`
	Image        string                       `json:"image" yaml:"image"`
	Traits       domain.TraitGroupDescription `json:"traits" yaml:"traits"`
	PathMap      map[string]string            `json:"path_map" yaml:"path_map"`
}

// nftool traits dump
func GenerateTraitGroupDescription(path string, out string) error {
	traitGroupDescription, pathMap, err := MakeTraitGroupDescriptionFromDir(path)
	if err != nil {
		return err
	}

	config := Config{
		Name:         "My NFT Name #{id}",
		Description:  "An awesome collection",
		ExternalLink: "https://myexternallink.com",
		Image:        "https://api.xyz.com/{id}.jpg",
		Traits:       traitGroupDescription,
		PathMap:      pathMap,
	}
	return utils.WriteFileAsYaml(config, out)
}

// nftool traits make
func GenerateTraitCollection(path string, out string, amount int) error {
	info, err := os.Stat(path)
	if os.IsNotExist(err) {
		return fmt.Errorf("path '%s' does not exist", path)
	}
	rand.Seed(time.Now().UnixNano())
	if info.IsDir() {
		return generateTraitCollectionFromDir(path, out, amount)
	}
	return generateTraitCollectionFromConfig(path, out, amount)
}

func generateTraitCollectionFromDir(path string, out string, amount int) error {
	traitCollection, err := MakeTraitCollectionFromDir(path, amount)
	if err != nil {
		return err
	}
	return utils.WriteFileAsJson(traitCollection, out)
}

func generateTraitCollectionFromConfig(path string, out string, amount int) error {
	traitCollection, err := MakeTraitCollectionFromConfigFile(path, amount)
	if err != nil {
		return err
	}
	return utils.WriteFileAsJson(traitCollection, out)
}

// nftool traits shuffle
func ShuffleCollection(path string) error {
	collection, err := LoadTraitCollectionFromFile(path)
	if err != nil {
		return err
	}

	rand.Seed(time.Now().UnixNano())
	collection.Shuffle()

	err = utils.WriteFileAsJson(collection, path)
	if err != nil {
		return err
	}

	return nil
}

// nftool traits merge
func MergeCollections(paths []string, out string) error {
	collections, err := loadMultipleTraitCollectionFiles(paths)
	if err != nil {
		return err
	}
	mergedCollections, _ := collections[0].Merge(collections[1:])
	err = utils.WriteFileAsJson(mergedCollections, out)
	if err != nil {
		return err
	}
	return nil
}

// nftool traits collection
func FindCollisions(files []string, out string) error {
	collections, err := loadMultipleTraitCollectionFiles(files)
	if err != nil {
		return err
	}

	collisions := collections[0].Collisions(collections[1:])

	err = utils.WriteFileAsJson(collisions, out)
	if err != nil {
		return err
	}

	return nil
}
