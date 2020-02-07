package matches

import (
	"time"

	"github.com/asalih/guardian/helpers"
)

//MatchResult Match result
type MatchResult struct {
	//	MatchedRules []*engine.Rule
	IsMatched bool
	StartTime time.Time
	Elapsed   int64
}

//NewMatchResult Inits match result
func NewMatchResult(isMatched bool) *MatchResult {
	return &MatchResult{isMatched, time.Now(), 0}
}

//SetMatch ...
func (m *MatchResult) SetMatch(isMatched bool) *MatchResult {
	m.IsMatched = isMatched

	m.Elapsed = helpers.CalcTime(m.StartTime, time.Now())

	return m
}
