package fs

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/alephao/nftool/traits/domain"
	"github.com/alephao/nftool/utils"
)

func LoadConfigFromFile(path string) (Config, error) {
	config := Config{}
	err := utils.LoadYamlFileIntoStruct(path, &config)
	if err != nil {
		return Config{}, err
	}
	return config, nil
}

func LoadTraitCollectionFromFile(path string) (domain.TraitCollection, error) {
	traits := domain.TraitCollection{}
	err := utils.LoadJsonFileIntoStruct(path, &traits)
	if err != nil {
		return nil, err
	}
	return traits, nil
}

func MakeTraitCollectionFromDir(dirPath string, amount int) (domain.TraitCollection, error) {
	traitGroupDescription, _, err := MakeTraitGroupDescriptionFromDir(dirPath)
	if err != nil {
		return nil, err
	}
	traitCollection, err := traitGroupDescription.RandomTraitCollection(amount)

	if err != nil {
		return nil, err
	}

	return traitCollection, nil
}

func MakeTraitCollectionFromConfigFile(path string, amount int) (domain.TraitCollection, error) {
	config, err := LoadConfigFromFile(path)
	if err != nil {
		return nil, err
	}

	traitCollection, err := config.Traits.RandomTraitCollection(amount)
	if err != nil {
		return nil, err
	}

	return traitCollection, nil
}

func MakeTraitGroupDescriptionFromDir(dirPath string) (domain.TraitGroupDescription, map[string]string, error) {
	dirs, err := utils.LsDirs(dirPath)
	if err != nil {
		return domain.TraitGroupDescription{}, nil, err
	}

	traitGroupDescription := domain.TraitGroupDescription{}
	pathMap := map[string]string{}

	for _, dir := range dirs {
		path := fmt.Sprintf("%s/%s", dirPath, dir)
		traitDescription, _pathMap, err := getTraitDescriptionFromDir(path)
		if err != nil {
			return domain.TraitGroupDescription{}, nil, err
		}

		for key, val := range _pathMap {
			pathMap[key] = val
		}

		traitGroupDescription = append(traitGroupDescription, traitDescription)
	}

	return traitGroupDescription, pathMap, nil
}

func getTraitDescriptionFromDir(dirPath string) (domain.TraitDescription, map[string]string, error) {
	traitName, err := getTraitName(dirPath)
	if err != nil {
		return domain.TraitDescription{}, nil, fmt.Errorf("failed to get trait name for path '%s': %s", dirPath, err.Error())
	}

	isOptional := strings.HasSuffix(traitName, "_")
	optionalWeight := 1
	traitName = strings.TrimSuffix(traitName, "_")

	variantDescriptions, paths, err := getTraitVariantDescriptionsFromFileSystem(dirPath)
	if err != nil {
		return domain.TraitDescription{}, nil, fmt.Errorf("failed to get trait variant descriptions for path '%s': %s", dirPath, err.Error())
	}

	pathMap := map[string]string{}

	for i, path := range paths {
		key := fmt.Sprintf("%s/%s", traitName, variantDescriptions[i].Value)
		value := path
		pathMap[key] = value
	}

	return domain.TraitDescription{
		TraitType:      traitName,
		Variants:       variantDescriptions,
		IsOptional:     isOptional,
		OptionalWeight: uint(optionalWeight),
	}, pathMap, nil
}

func getTraitName(dirPath string) (string, error) {
	components := strings.Split(dirPath, "-")

	if len(components) == 2 {
		return components[1], nil
	}

	return "", fmt.Errorf("invalid dir name '%s'", dirPath)
}

func getTraitVariantDescriptionsFromFileSystem(path string) ([]domain.TraitVariantDescription, []string, error) {
	files, err := utils.LsFiles(path)

	if err != nil {
		return nil, nil, err
	}

	traitVariantDescriptions := []domain.TraitVariantDescription{}
	paths := []string{}

	for _, fileName := range files {
		traitVariantDescription, err := getTraitVariantDescriptionFromFileName(fileName)
		if err != nil {
			return nil, nil, err
		}

		traitPath := fmt.Sprintf("%s/%s", path, fileName)
		traitVariantDescriptions = append(traitVariantDescriptions, traitVariantDescription)
		paths = append(paths, traitPath)
	}

	return traitVariantDescriptions, paths, nil
}

// RarityWeight-Name of Variant.png
// E.g.: 2-My Background.png
// -> Name: My Background
// -> Weight: 2
func getTraitVariantDescriptionFromFileName(name string) (domain.TraitVariantDescription, error) {
	nameWithNoExtension := strings.Split(name, ".")[0]
	parts := strings.Split(nameWithNoExtension, "-")

	if len(parts) >= 2 {
		rarity, err := strconv.Atoi(parts[0])
		if err != nil {
			return domain.TraitVariantDescription{}, fmt.Errorf("invalid fileName '%s': %s", name, err.Error())
		}

		weight := uint(rarity)
		if weight == 0 {
			return domain.TraitVariantDescription{}, fmt.Errorf("invalid rarirty '%d' in file '%s': %s", rarity, name, err.Error())
		}

		return domain.TraitVariantDescription{
			Value:  strings.Join(parts[1:], "-"),
			Weight: weight,
		}, nil
	}

	return domain.TraitVariantDescription{}, fmt.Errorf("invalid fileName '%s'", name)
}
