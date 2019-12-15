package models

import "time"

//MatchResult Match result
type MatchResult struct {
	MatchedPayloads []*PayloadData
	Elapsed         int64
	IsMatched       bool
}

//FirewallMatchResult Firewall rules match result
type FirewallMatchResult struct {
	FirewallRule *FirewallRule
	Elapsed      int64
	IsMatched    bool
}

//NewMatchResult Inits match result
func NewMatchResult(isMatched bool) *MatchResult {
	return &MatchResult{nil, 0, isMatched}
}

//NewFirewallMatchResult Inits fw match result
func NewFirewallMatchResult(rule *FirewallRule, isMatched bool) *FirewallMatchResult {
	return &FirewallMatchResult{rule, 0, isMatched}
}

//Time calculates elapsed time
func (m *MatchResult) Time(now time.Time) *MatchResult {
	m.Elapsed = CalcTime(now)

	return m
}

//Append ...
func (m *MatchResult) Append(payload *PayloadData) *MatchResult {
	m.MatchedPayloads = append(m.MatchedPayloads, payload)

	return m
}

//SetMatch ...
func (m *MatchResult) SetMatch(isMatched bool) *MatchResult {
	m.IsMatched = isMatched

	return m
}

//Time calculates elapsed time
func (m *FirewallMatchResult) Time(now time.Time) *FirewallMatchResult {
	m.Elapsed = CalcTime(now)

	return m
}
