package domain

type ProvenanceReport struct {
	FinalProofHash     string       `json:"final_proof_hash"`
	ConcatenatedHashes string       `json:"concatenated_hashes"`
	StartingIndex      int          `json:"starting_index"`
	Hashes             []HashOfFile `json:"hashes"`
}

type HashOfFile struct {
	FileName             string `json:"file_name"`
	InitialSequenceIndex int    `json:"index"`
	TokenId              int    `json:"token_id"`
	Total                int    `json:"total"`
	Hash                 string `json:"hash"`
}
