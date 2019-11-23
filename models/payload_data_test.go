package models

import (
	"testing"
)

func TestInitRequestPayloadDataCollectionFile(t *testing.T) {
	tests := []struct {
		name string
		path string
	}{
		{"InitRequestPayload", "../requestPayloads.json"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InitRequestPayloadDataCollectionFile(tt.path)
		})
	}
}

func TestInitResponsePayloadDataCollectionFile(t *testing.T) {
	tests := []struct {
		name string
		path string
	}{
		{"InitResponsePayload", "../responsePayloads.json"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InitResponsePayloadDataCollectionFile(tt.path)
		})
	}
}
