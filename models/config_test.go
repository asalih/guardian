package models

import "testing"

func TestInitConfigFile(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"ConfigInit"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InitConfigFile("../appsettings")
		})
	}
}
