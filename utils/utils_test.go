package utils

import "testing"

func TestHash(t *testing.T) {
	a := []string{"A", "B", "C"}
	b := []string{"A", "B", "C"}

	aHash := Hash(a)
	bHash := Hash(b)

	if aHash != bHash {
		t.Errorf("Expected '%x' to be equal to '%x'", aHash, bHash)
	}
}
