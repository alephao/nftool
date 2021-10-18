package domain

import (
	"fmt"

	wr "github.com/mroth/weightedrand"
)

const TraitVariantDescriptionNoneValue = "none"

type TraitVariantDescription struct {
	Value  string `json:"value"`
	Weight uint   `json:"weight"`
}

func MakeVariantDescriptionNone(weight uint) TraitVariantDescription {
	return TraitVariantDescription{
		Value:  TraitVariantDescriptionNoneValue,
		Weight: weight,
	}
}

type TraitDescription struct {
	TraitType      string                    `json:"trait_type" yaml:"trait_type"`
	Variants       []TraitVariantDescription `json:"variants"`
	IsOptional     bool                      `json:"is_optional" yaml:"is_optional"`
	OptionalWeight uint                      `json:"optional_weight" yaml:"optional_weight"`
}

func (desc TraitDescription) RandomTraitInstance() (TraitInstance, error) {
	chooser, err := desc.makeVariantChooser()
	if err != nil {
		return TraitInstance{}, fmt.Errorf("failed to make variant chooser: %w", err)
	}
	variantDescription := chooser.Pick().(TraitVariantDescription)
	return TraitInstance{
		TraitType: desc.TraitType,
		Value:     variantDescription.Value,
	}, nil
}

func (desc TraitDescription) makeVariantChooser() (*wr.Chooser, error) {
	choices := []wr.Choice{}

	for _, variant := range desc.Variants {
		choice := wr.NewChoice(variant, variant.Weight)
		choices = append(choices, choice)
	}

	if desc.IsOptional && desc.OptionalWeight > 0 {
		variantNone := MakeVariantDescriptionNone(desc.OptionalWeight)
		choice := wr.NewChoice(variantNone, desc.OptionalWeight)
		choices = append(choices, choice)
	}

	return wr.NewChooser(choices...)
}

type TraitGroupDescription []TraitDescription

func (c TraitGroupDescription) RandomTraitInstances() (TraitGroup, error) {
	traitGroup := TraitGroup{}
	for _, traitDescription := range c {
		traitInstance, err := traitDescription.RandomTraitInstance()
		if err != nil {
			return nil, err
		}
		traitGroup = append(traitGroup, traitInstance)
	}
	return traitGroup, nil
}

func (c TraitGroupDescription) RandomTraitCollection(amount int) (TraitCollection, error) {
	traitGroups := TraitCollection{}
	hashes := map[string]bool{}
	count := 0
	for count < amount {
		randomInstances, err := c.RandomTraitInstances()
		if err != nil {
			return nil, err
		}
		instanceHash := hash(randomInstances)

		if _, ok := hashes[instanceHash]; ok {
			continue
		}

		hashes[instanceHash] = true
		traitGroups = append(traitGroups, randomInstances)
		count++
	}
	return traitGroups, nil
}
