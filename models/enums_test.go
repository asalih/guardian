package models

import (
	"reflect"
	"testing"
)

func TestWafAction_ToString(t *testing.T) {
	tests := []struct {
		name   string
		action WafAction
		want   string
	}{
		{"allow", WafActionAllow, "Allow"},
		{"block", WafActionBlock, "Block"},
		{"log", WafActionLog, "Log"},
		{"remove", WafActionRemove, "Remove"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.action.ToString(); got != tt.want {
				t.Errorf("WafAction.ToString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetWafAction(t *testing.T) {
	type args struct {
		action string
	}
	tests := []struct {
		name   string
		action string
		want   WafAction
	}{
		{"block", "block", WafActionBlock},
		{"allow", "allow", WafActionAllow},
		{"log", "log", WafActionLog},
		{"remove", "remove", WafActionRemove},
		{"none", "none", 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetWafAction(tt.action); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetWafAction() = %v, want %v", got, tt.want)
			}
		})
	}
}
