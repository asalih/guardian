package models

//DisruptiveAction WAF Disruptive action
type DisruptiveAction uint8

const (
	//DisruptiveActionPass Pass
	DisruptiveActionPass DisruptiveAction = iota
	//DisruptiveActionBlock Blocks
	DisruptiveActionBlock
	//DisruptiveActionDrop Drop
	DisruptiveActionDrop
	//DisruptiveActionDeny Deny
	DisruptiveActionDeny
	//DisruptiveActionProxy Proxy(WTF)
	DisruptiveActionProxy
)

//LogAction Log action
type LogAction uint8

const (
	//LogActionLog Log
	LogActionLog LogAction = iota
	//LogActionNoLog No log
	LogActionNoLog
)

//Phase WAF Rule check phase
type Phase uint8

const (
	//Phase1 First
	Phase1 Phase = iota
	//Phase2 Second
	Phase2
	//Phase3 Third
	Phase3
	//Phase4 Fourth
	Phase4
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
