package matches

import (
	"time"

	"github.com/asalih/guardian/helpers"
)

//FirewallMatchResult Firewall rules match result
type FirewallMatchResult struct {
	IsMatched bool
	EndTime   time.Time
}

//NewFirewallMatchResult Inits fw match result
func NewFirewallMatchResult( /*rule *FirewallRule,*/ isMatched bool) *FirewallMatchResult {
	return &FirewallMatchResult{ /*rule,*/ isMatched, time.Time{}}
}

func (m *FirewallMatchResult) Elapsed(start time.Time) int64 {
	if !m.IsMatched {
		return 0
	}

	return helpers.CalcTime(start, m.EndTime)
}
