package domain

import (
	traits "github.com/alephao/nftool/traits/domain"
)

type TraitsRarityReport struct {
	Total  int              `json:"total"`
	Traits TraitsRarityData `json:"traits"`
}

type TraitRarityData struct {
	Total      int
	Percentage float64
}

type TraitsRarityData = map[string]map[string]TraitRarityData

func MakeTraitRarityReport(traits []traits.TraitGroup) TraitsRarityReport {
	n := len(traits)
	rawData := TraitsRarityData{}

	i := 0
	for i < n {
		for _, attr := range traits[i] {

			_, traitExist := rawData[attr.TraitType]

			if !traitExist {
				rawData[attr.TraitType] = map[string]TraitRarityData{
					attr.Value: {
						Total:      0,
						Percentage: 0,
					},
				}
			}

			val := rawData[attr.TraitType][attr.Value]
			rawData[attr.TraitType][attr.Value] = TraitRarityData{
				Total:      val.Total + 1,
				Percentage: 0,
			}
		}
		i++
	}

	for key, val := range rawData {
		for keyData, valData := range val {
			total := valData.Total
			rawData[key][keyData] = TraitRarityData{
				Total:      total,
				Percentage: float64(total) / float64(n),
			}
		}
	}

	report := TraitsRarityReport{
		Total:  n,
		Traits: rawData,
	}

	return report
}
