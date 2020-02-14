package models

import (
	"testing"
)

func TestDisruptiveAction_ToString(t *testing.T) {
	tests := []struct {
		name   string
		action DisruptiveAction
		want   string
	}{
		{"pass", DisruptiveActionPass, "pass"},
		{"block", DisruptiveActionBlock, "block"},
		{"drop", DisruptiveActionDrop, "drop"},
		{"deny", DisruptiveActionDeny, "deny"},
		{"proxy", DisruptiveActionProxy, "proxy"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.action.ToString(); got != tt.want {
				t.Errorf("DisruptiveAction.ToString() = %v, want %v", got, tt.want)
			}
		})
	}
}
