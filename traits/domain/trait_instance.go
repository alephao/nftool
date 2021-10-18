package domain

type TraitInstance struct {
	Value     string `json:"value"`
	TraitType string `json:"trait_type"`
}

func MakeTraitInstance(value, traitType string) TraitInstance {
	return TraitInstance{
		Value:     value,
		TraitType: traitType,
	}
}

type TraitGroup []TraitInstance
