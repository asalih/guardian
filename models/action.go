package models

//Action definition for Rules
type Action struct {
	ID               string
	Phase            int
	DisruptiveAction DisruptiveAction
	LogAction        LogAction
}
