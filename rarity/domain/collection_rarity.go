package domain

import (
	"sort"

	traits "github.com/alephao/nftool/traits/domain"
)

// Collection Rarity

type CollectionItemRarity struct {
	Index int
	Score float64
	Attrs traits.TraitGroup
}

func MakeCollectionRarityReport(collection traits.TraitCollection) []CollectionItemRarity {
	traitRarityReport := MakeTraitRarityReport(collection)
	collectionRarity := []CollectionItemRarity{}
	for i, attrs := range collection {
		score := CalculateRarityScore(attrs, traitRarityReport)
		rarity := CollectionItemRarity{
			Index: i,
			Score: score,
			Attrs: attrs,
		}
		collectionRarity = append(collectionRarity, rarity)
	}
	sort.Slice(collectionRarity, func(i, j int) bool {
		return collectionRarity[i].Score < collectionRarity[j].Score
	})
	return collectionRarity
}

func CalculateRarityScore(group traits.TraitGroup, report TraitsRarityReport) float64 {
	score := 1.0
	for _, attr := range group {
		score *= report.Traits[attr.TraitType][attr.Value].Percentage
	}
	return score
}
