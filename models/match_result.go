package models

import "time"

//MatchResult Match result
type MatchResult struct {
	MatchedPayload *PayloadData
	Elapsed        int64
	IsMatched      bool
}

//FirewallMatchResult Firewall rules match result
type FirewallMatchResult struct {
	FirewallRule *FirewallRule
	Elapsed      int64
	IsMatched    bool
}

//NewMatchResult Inits match result
func NewMatchResult(payload *PayloadData, isMatched bool) *MatchResult {
	return &MatchResult{payload, 0, isMatched}
}

//NewFirewallMatchResult Inits fw match result
func NewFirewallMatchResult(rule *FirewallRule, isMatched bool) *FirewallMatchResult {
	return &FirewallMatchResult{rule, 0, isMatched}
}

//Time calculates elapsed time
func (m MatchResult) Time(now time.Time) *MatchResult {
	m.Elapsed = CalcTime(now)

	return &m
}

//Time calculates elapsed time
func (m FirewallMatchResult) Time(now time.Time) *FirewallMatchResult {
	m.Elapsed = CalcTime(now)

	return &m
}
