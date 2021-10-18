package domain

import (
	"math/rand"
	"testing"

	"github.com/alephao/nftool/utils"
)

func TestTraitCollectionShuffle(t *testing.T) {
	t.Run("simple description", func(t *testing.T) {
		collection := TraitCollection{
			TraitGroup{
				TraitInstance{
					TraitType: "test",
					Value:     "0",
				},
			},
			TraitGroup{
				TraitInstance{
					TraitType: "test",
					Value:     "1",
				},
			},
		}

		rand.Seed(2)
		collection.Shuffle()

		expected := TraitCollection{
			TraitGroup{
				TraitInstance{
					TraitType: "test",
					Value:     "1",
				},
			},
			TraitGroup{
				TraitInstance{
					TraitType: "test",
					Value:     "0",
				},
			},
		}

		assertInterfaceEqual(t, expected, collection)
	})
}

func TestTraitCollectionMerge(t *testing.T) {
	collections := []TraitCollection{
		{
			{
				MakeTraitInstance("Value 1", "Trait 1"),
				MakeTraitInstance("Value 2", "Trait 2"),
			},
			{
				MakeTraitInstance("Value 2", "Trait 1"),
				MakeTraitInstance("Value 1", "Trait 2"),
			},
			{
				MakeTraitInstance("Value 2", "Trait 1"),
				MakeTraitInstance("none", "Trait 2"),
			},
		},
		{
			{
				MakeTraitInstance("Value 2", "Trait 1"),
				MakeTraitInstance("Value 2", "Trait 2"),
			},
			{
				MakeTraitInstance("Value 1", "Trait 1"),
				MakeTraitInstance("Value 1", "Trait 2"),
			},
			{
				MakeTraitInstance("Value 1", "Trait 1"),
				MakeTraitInstance("Value 2", "Trait 2"),
			},
		},
		{
			{
				MakeTraitInstance("Value 2", "Trait 1"),
				MakeTraitInstance("Value 3", "Trait 2"),
			},
			{
				MakeTraitInstance("Value 3", "Trait 1"),
				MakeTraitInstance("Value 3", "Trait 2"),
			},
			{
				MakeTraitInstance("Value 1", "Trait 1"),
				MakeTraitInstance("Value 2", "Trait 2"),
			},
		},
	}

	got, removedCollisions := collections[0].Merge(collections[1:])

	if removedCollisions != 2 {
		t.Errorf("Expected %d removed collisions, got %d", 2, removedCollisions)
	}

	expected := TraitCollection{
		{
			MakeTraitInstance("Value 1", "Trait 1"),
			MakeTraitInstance("Value 2", "Trait 2"),
		},
		{
			MakeTraitInstance("Value 2", "Trait 1"),
			MakeTraitInstance("Value 1", "Trait 2"),
		},
		{
			MakeTraitInstance("Value 2", "Trait 1"),
			MakeTraitInstance("none", "Trait 2"),
		},
		{
			MakeTraitInstance("Value 2", "Trait 1"),
			MakeTraitInstance("Value 2", "Trait 2"),
		},
		{
			MakeTraitInstance("Value 1", "Trait 1"),
			MakeTraitInstance("Value 1", "Trait 2"),
		},
		{
			MakeTraitInstance("Value 2", "Trait 1"),
			MakeTraitInstance("Value 3", "Trait 2"),
		},
		{
			MakeTraitInstance("Value 3", "Trait 1"),
			MakeTraitInstance("Value 3", "Trait 2"),
		},
	}

	assertInterfaceEqual(t, expected, got)
}

func TestTraitCollectionCollision(t *testing.T) {
	collections := []TraitCollection{
		{
			{
				MakeTraitInstance("Value 1", "Trait 1"),
				MakeTraitInstance("Value 2", "Trait 2"),
			},
			{
				MakeTraitInstance("Value 2", "Trait 1"),
				MakeTraitInstance("Value 1", "Trait 2"),
			},
			{
				MakeTraitInstance("Value 2", "Trait 1"),
				MakeTraitInstance("none", "Trait 2"),
			},
		},
		{
			{
				MakeTraitInstance("Value 2", "Trait 1"),
				MakeTraitInstance("Value 2", "Trait 2"),
			},
			{
				MakeTraitInstance("Value 1", "Trait 1"),
				MakeTraitInstance("Value 1", "Trait 2"),
			},
			{
				MakeTraitInstance("Value 1", "Trait 1"),
				MakeTraitInstance("Value 2", "Trait 2"),
			},
		},
		{
			{
				MakeTraitInstance("Value 2", "Trait 1"),
				MakeTraitInstance("Value 3", "Trait 2"),
			},
			{
				MakeTraitInstance("Value 3", "Trait 1"),
				MakeTraitInstance("Value 3", "Trait 2"),
			},
			{
				MakeTraitInstance("Value 1", "Trait 1"),
				MakeTraitInstance("Value 2", "Trait 2"),
			},
		},
	}

	got := collections[0].Collisions(collections[1:])

	group := TraitGroup{
		MakeTraitInstance("Value 1", "Trait 1"),
		MakeTraitInstance("Value 2", "Trait 2"),
	}
	hash := utils.Hash(group)
	expected := []CollectionItemCollision{
		{
			Hash:  string(hash[:]),
			Group: group,
			Where: map[int]int{
				0: 0,
				1: 2,
				2: 2,
			},
		},
	}

	assertInterfaceEqual(t, expected, got)
}
