package models

//MatchResult Match result
type MatchResult struct {
	MatchedPayload *PayloadData
	IsMatched      bool
}

//NewMatchResult Inits match result
func NewMatchResult(payload *PayloadData, isMatched bool) *MatchResult {
	return &MatchResult{payload, isMatched}
}
