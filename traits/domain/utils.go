package domain

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
)

func hash(v interface{}) string {
	j, _ := json.Marshal(v)
	hash := sha256.Sum256(j)
	return hex.EncodeToString(hash[:])
}
