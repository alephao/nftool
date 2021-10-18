package domain

import (
	"math/rand"
	"testing"
)

func TestRandomTraitInstance(t *testing.T) {
	t.Run("simple description", func(t *testing.T) {
		rand.Seed(0)

		traitDescription := TraitDescription{
			TraitType: "Test Trait",
			Variants: []TraitVariantDescription{
				{Value: "A", Weight: 1},
				{Value: "B", Weight: 1},
			},
			IsOptional:     false,
			OptionalWeight: 0,
		}

		got, err := traitDescription.RandomTraitInstance()
		assertNoError(t, err)
		expected := TraitInstance{
			Value:     "A",
			TraitType: "Test Trait",
		}
		assertInterfaceEqual(t, expected, got)
	})

	t.Run("optional description", func(t *testing.T) {
		rand.Seed(0)

		traitDescription := TraitDescription{
			TraitType: "Test Trait",
			Variants: []TraitVariantDescription{
				{Value: "A", Weight: 1},
				{Value: "B", Weight: 1},
			},
			IsOptional:     true,
			OptionalWeight: 2,
		}

		got, err := traitDescription.RandomTraitInstance()
		assertNoError(t, err)
		expected := TraitInstance{
			Value:     "none",
			TraitType: "Test Trait",
		}
		assertInterfaceEqual(t, expected, got)
	})
}

func TestRandomTraitInstances(t *testing.T) {
	t.Run("simple trait group", func(t *testing.T) {
		rand.Seed(0)

		traitGroupDescription := TraitGroupDescription{
			TraitDescription{
				TraitType: "0",
				Variants: []TraitVariantDescription{
					{Value: "A", Weight: 1},
					{Value: "B", Weight: 1},
				},
				IsOptional:     false,
				OptionalWeight: 0,
			},
			TraitDescription{
				TraitType: "1",
				Variants: []TraitVariantDescription{
					{Value: "A", Weight: 1},
					{Value: "B", Weight: 1},
				},
				IsOptional:     false,
				OptionalWeight: 0,
			},
		}

		got, err := traitGroupDescription.RandomTraitInstances()
		assertNoError(t, err)
		expected := TraitGroup{
			TraitInstance{
				Value:     "A",
				TraitType: "0",
			},
			TraitInstance{
				Value:     "A",
				TraitType: "1",
			},
		}
		assertInterfaceEqual(t, expected, got)
	})

	t.Run("trait group with optional", func(t *testing.T) {
		rand.Seed(0)

		traitGroupDescription := TraitGroupDescription{
			TraitDescription{
				TraitType: "0",
				Variants: []TraitVariantDescription{
					{Value: "A", Weight: 1},
					{Value: "B", Weight: 1},
				},
				IsOptional:     false,
				OptionalWeight: 0,
			},
			TraitDescription{
				TraitType: "1",
				Variants: []TraitVariantDescription{
					{Value: "A", Weight: 1},
					{Value: "B", Weight: 1},
				},
				IsOptional:     true,
				OptionalWeight: 2,
			},
		}

		got, err := traitGroupDescription.RandomTraitInstances()
		assertNoError(t, err)
		expected := TraitGroup{
			TraitInstance{
				Value:     "A",
				TraitType: "0",
			},
			TraitInstance{
				Value:     "none",
				TraitType: "1",
			},
		}
		assertInterfaceEqual(t, expected, got)
	})
}
