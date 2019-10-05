package models

type WafAction int

const (
	WafActionBlock = iota
	WafActionAllow
	WafActionRemove
	WafActionLog
)

const (
	LogTypeWAF = iota
	LogTypeFirewall
)

//ToString for waf action
func (action WafAction) ToString() string {
	return [...]string{"Block", "Allow", "Remove", "Log"}[action]
}

//GetWafAction Gets the waf action with given action string
func GetWafAction(action string) WafAction {
	switch action {
	case "block":
		return 0
	case "allow":
		return 1
	case "remove":
		return 2
	case "log":
		return 3
	}

	return 0
}
