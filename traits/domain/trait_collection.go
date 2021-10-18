package domain

import (
	"math/rand"

	"github.com/alephao/nftool/utils"
)

type TraitCollection []TraitGroup

func (collection TraitCollection) Shuffle() {
	rand.Shuffle(len(collection), func(i, j int) {
		collection[i], collection[j] = collection[j], collection[i]
	})
}

func (collection TraitCollection) Merge(others []TraitCollection) (TraitCollection, int) {
	alreadyAddedAttribute := map[[32]byte]bool{}
	removedCollisions := 0

	mergedCollection := TraitCollection{}

	for _, traits := range collection {
		mergedCollection = append(mergedCollection, traits)
		hash := utils.Hash(traits)
		alreadyAddedAttribute[hash] = true
	}

	for _, otherCollection := range others {
		for _, traits := range otherCollection {
			hash := utils.Hash(traits)
			_, ok := alreadyAddedAttribute[hash]
			if !ok {
				mergedCollection = append(mergedCollection, traits)
				alreadyAddedAttribute[hash] = true
			} else {
				removedCollisions++
			}
		}
	}

	return mergedCollection, removedCollisions
}

type CollectionItemCollision struct {
	Hash  string
	Group TraitGroup
	Where map[int]int
}

func (collection TraitCollection) Collisions(others []TraitCollection) []CollectionItemCollision {
	// {
	//   hash_a: {
	//     file_index: array_index
	//   }
	// }
	//
	// E.g.:
	// {
	//   hash_a: {
	//     2: 4
	//   }
	// }
	hashToPosition := map[[32]byte]map[int]int{}

	collectionsToCompare := []TraitCollection{collection}
	collectionsToCompare = append(collectionsToCompare, others...)

	for i, attrs := range collectionsToCompare {
		for j, cia := range attrs {
			hash := utils.Hash(cia)
			_, ok := hashToPosition[hash]
			if !ok {
				hashToPosition[hash] = map[int]int{}
			}
			hashToPosition[hash][i] = j
		}
	}

	collisions := []CollectionItemCollision{}
	for key, val := range hashToPosition {
		nOfKeys := len(val)
		if nOfKeys > 1 {
			var collisionGroup TraitGroup
			for filePos, arrPos := range val {
				collisionGroup = collectionsToCompare[filePos][arrPos]
				break
			}

			collision := CollectionItemCollision{
				Hash:  string(key[:]),
				Group: collisionGroup,
				Where: val,
			}
			collisions = append(collisions, collision)
		}
	}

	return collisions
}
