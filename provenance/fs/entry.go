package fs

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"

	"github.com/alephao/nftool/provenance/domain"
	"github.com/alephao/nftool/utils"
)

func ProvenanceReportFromDir(dirPath, out string, startingIndex int) error {
	hashes, err := HashOfImagesInDir(dirPath, startingIndex)
	if err != nil {
		return err
	}

	hashesOnly := []string{}
	for _, hashOfFile := range hashes {
		hashesOnly = append(hashesOnly, hashOfFile.Hash)
	}
	concatenatedHashes := strings.Join(hashesOnly, "")
	finalProof, err := Sha256String(concatenatedHashes)
	if err != nil {
		return err
	}

	report := domain.ProvenanceReport{
		FinalProofHash:     finalProof,
		ConcatenatedHashes: concatenatedHashes,
		StartingIndex:      startingIndex,
		Hashes:             hashes,
	}

	return utils.WriteFileAsJson(report, out)
}

func HashOfImagesInDir(dirPath string, startingIndex int) ([]domain.HashOfFile, error) {
	imgs, err := utils.LsFiles(dirPath)
	if err != nil {
		return nil, err
	}

	total := len(imgs)
	hashes := []domain.HashOfFile{}
	for _, img := range imgs {
		hash, err := HashOfImage(fmt.Sprintf("%s/%s", dirPath, img))
		if err != nil {
			return nil, err
		}
		tokenId, err := strconv.Atoi(strings.Split(img, ".")[0])
		if err != nil {
			return nil, err
		}
		initialSeqIndex := (tokenId + startingIndex) % total
		hashOfFile := domain.HashOfFile{
			FileName:             img,
			InitialSequenceIndex: initialSeqIndex,
			TokenId:              tokenId,
			Total:                total,
			Hash:                 hash,
		}
		hashes = append(hashes, hashOfFile)
	}

	sort.Slice(hashes, func(i, j int) bool {
		return hashes[i].InitialSequenceIndex < hashes[j].InitialSequenceIndex
	})

	return hashes, nil
}

func HashOfImage(path string) (string, error) {
	hasher := sha256.New()
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	hasher.Write(file)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(hasher.Sum(nil)), nil
}

func Sha256String(str string) (string, error) {
	hasher := sha256.New()
	_, err := hasher.Write([]byte(str))
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(hasher.Sum(nil)), nil
}
