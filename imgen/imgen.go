package imgen

import (
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"os"
	"runtime"
	"sync"

	traits "github.com/alephao/nftool/traits/domain"
	traits_fs "github.com/alephao/nftool/traits/fs"
	"github.com/alephao/nftool/utils"
)

func GenerateImagesFromCollectionAttributesJson(width, height, startingIndex int, collectionPath, configPath, outDir string, parallel, saveAsPng, onlyMissing, transparent bool) error {
	var config traits_fs.Config
	err := utils.LoadYamlFileIntoStruct(configPath, &config)
	if err != nil {
		return err
	}

	collectionAttributes, err := traits_fs.LoadTraitCollectionFromFile(collectionPath)
	if err != nil {
		return err
	}

	if onlyMissing {
		return GenerateMissingImagesFromCollectionAttributes(width, height, collectionAttributes, config.PathMap, outDir, saveAsPng, transparent)
	}

	if parallel {
		return GenerateManyImagesFromCollectionAttributesParallel(width, height, startingIndex, collectionAttributes, config.PathMap, outDir, saveAsPng, transparent)
	} else {
		return GenerateManyImagesFromCollectionAttributes(width, height, startingIndex, collectionAttributes, config.PathMap, outDir, saveAsPng, transparent)
	}
}

func GenerateManyImagesFromCollectionAttributesParallel(width, height, startingIndex int, collectionAttributes []traits.TraitGroup, layersMap map[string]string, outputDir string, saveAsPng, transparent bool) error {
	var chunks []traits.TraitCollection
	numCpu := runtime.NumCPU()
	chunkSize := (len(collectionAttributes) + numCpu - 1) / numCpu

	for i := 0; i < len(collectionAttributes); i += chunkSize {
		end := i + chunkSize

		if end > len(collectionAttributes) {
			end = len(collectionAttributes)
		}

		chunks = append(chunks, collectionAttributes[i:end])
	}

	var wg sync.WaitGroup

	for i, chunk := range chunks {
		wg.Add(1)
		startingIndex := i * chunkSize
		chunkCopy := make(traits.TraitCollection, chunkSize)
		copy(chunkCopy, chunk)
		go func() {
			GenerateManyImagesFromCollectionAttributes(width, height, startingIndex, chunkCopy, layersMap, outputDir, saveAsPng, transparent)
			wg.Done()
		}()
	}

	wg.Wait()

	return nil
}

func GenerateManyImagesFromCollectionAttributes(width, height, startingIndex int, collectionAttributes []traits.TraitGroup, layersMap map[string]string, outputDir string, saveAsPng, transparent bool) error {
	for i, traitGroup := range collectionAttributes {
		var extension string
		if saveAsPng {
			extension = "png"
		} else {
			extension = "jpg"
		}
		out := fmt.Sprintf("%s/%d.%s", outputDir, startingIndex+i, extension)
		if i%50 == 0 {
			fmt.Printf("Generating %s\n", out)
		}
		err := GenerateImageFromAttributes(width, height, traitGroup, layersMap, out, saveAsPng, transparent)
		if err != nil {
			// return err
			fmt.Printf("Failed to generate: %s\nError: %s\n", out, err.Error())
		}
	}
	return nil
}

func GenerateMissingImagesFromCollectionAttributes(width, height int, collection traits.TraitCollection, layersMap map[string]string, outputDir string, saveAsPng, transparent bool) error {
	missingIndexes, err := FindMissingIndexes(0, len(collection), outputDir, saveAsPng)
	if err != nil {
		return fmt.Errorf("failed to find missing indexes: %s", err.Error())
	}

	var extension string
	if saveAsPng {
		extension = "png"
	} else {
		extension = "jpg"
	}

	for _, missingIndex := range missingIndexes {
		out := fmt.Sprintf("%s/%d.%s", outputDir, missingIndex, extension)
		traitGroup := collection[missingIndex]
		fmt.Printf("Generating %s\n", out)
		err := GenerateImageFromAttributes(width, height, traitGroup, layersMap, out, saveAsPng, transparent)
		if err != nil {
			fmt.Printf("Failed to generate: %s\nError: %s\n", out, err.Error())
		}
	}
	return nil
}

func GenerateImageFromAttributes(width, height int, attributes traits.TraitGroup, layersMap map[string]string, outputPath string, saveAsPng, transparent bool) error {
	keys := []string{}
	layersPaths := []string{}
	for _, attr := range attributes {
		if attr.Value == "none" || attr.TraitType == "Special" {
			continue
		}
		key := fmt.Sprintf("%s/%s", attr.TraitType, attr.Value)
		keys = append(keys, key)
		layersPaths = append(layersPaths, layersMap[key])
	}
	return GenerateImageFromLayers(width, height, keys, layersPaths, outputPath, saveAsPng, transparent)
}

func GenerateImageFromLayers(width, height int, keys []string, layerPaths []string, outputPath string, saveAsPng, transparent bool) error {
	newImage := image.NewNRGBA(image.Rect(0, 0, width, height))

	for i, path := range layerPaths {
		imageFile, err := os.Open(path)
		if err != nil {
			return fmt.Errorf("failed to open image at path '%s'.\nkey: %s\nerr: %s", path, keys[i], err.Error())
		}

		defer imageFile.Close()

		img, err := png.Decode(imageFile)
		if err != nil {
			return fmt.Errorf("failed to decode image as png at path '%s': %s", path, err.Error())
		}

		if i == 0 && !transparent {
			draw.Draw(newImage, newImage.Bounds(), img, image.Point{}, draw.Src)
		} else {
			draw.Draw(newImage, newImage.Bounds(), img, image.Point{}, draw.Over)
		}
	}

	resultImg, err := os.Create(outputPath)
	if err != nil {
		return fmt.Errorf("failed to save image at path '%s': %s", outputPath, err.Error())
	}

	defer resultImg.Close()

	if saveAsPng {
		encoder := png.Encoder{
			CompressionLevel: png.NoCompression,
		}
		encoder.Encode(resultImg, newImage)
	} else {
		jpeg.Encode(resultImg, newImage, &jpeg.Options{
			Quality: jpeg.DefaultQuality,
		})
	}

	return nil
}

func FindMissingIndexes(from, to int, dir string, png bool) ([]int, error) {
	if to < from {
		return nil, fmt.Errorf("%d (from) should not be greater %d (to)", from, to)
	}

	var extension string
	if png {
		extension = "png"
	} else {
		extension = "jpg"
	}

	filesMissing := []int{}
	i := from
	for i <= to {
		path := fmt.Sprintf("%s/%d.%s", dir, i, extension)
		if !fileExists(path) {
			filesMissing = append(filesMissing, i)
		}
		i++
	}

	return filesMissing, nil
}

func fileExists(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	} else if os.IsNotExist(err) {
		return false
	} else {
		// Schrodinger, but return false anyway
		return false
	}
}
