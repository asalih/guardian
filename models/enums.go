package models

type DisruptiveAction uint16

const (
	DisruptiveActionPass DisruptiveAction = iota
	DisruptiveActionBlock
	DisruptiveActionDrop
	DisruptiveActionDeny
	DisruptiveActionProxy
)

type LogAction uint16

const (
	LogActionLog LogAction = iota
	LogActionNoLog
)

//ToString for waf action
func (action DisruptiveAction) ToString() string {
	return [...]string{"pass", "block", "drop", "deny", "proxy"}[action]
}

//GetDisruptiveAction Gets the waf action with given action string
func GetDisruptiveAction(action string) DisruptiveAction {
	switch action {
	case "pass":
		return DisruptiveActionPass
	case "block":
		return DisruptiveActionBlock
	case "drop":
		return DisruptiveActionDrop
	case "deny":
		return DisruptiveActionDeny
	case "proxy":
		return DisruptiveActionProxy
	}

	return DisruptiveActionBlock
}
