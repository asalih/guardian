package models

import (
	"net/http"
	"testing"
	"time"
)

func TestNewHTTPLog(t *testing.T) {
	tests := []struct {
		name string
		want *HTTPLog
	}{
		{"Init", &HTTPLog{"ID", "example.com", 0, 0, 0, 0, 0, 0, time.Now()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewHTTPLog()
			got.TargetID = "ID"
			got.RequestURI = "example.com"

			if got.TargetID != tt.want.TargetID || got.RequestURI != tt.want.RequestURI {
				t.Errorf("NewHTTPLog() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHTTPLog_Build(t *testing.T) {
	type args struct {
		target   *Target
		request  *http.Request
		response *http.Response
	}

	refHTTP := NewHTTPLog()
	refHTTP.TargetID = "ID"
	refHTTP.RequestURI = "www.netsparker.com"

	refRequest, _ := http.NewRequest("GET", "www.netsparker.com", nil)
	refRequest.RequestURI = "www.netsparker.com"
	target := new(Target)
	target.ID = "ID"
	target.Domain = "www.netsparker.com"

	tests := []struct {
		name string
		h    HTTPLog
		args args
		want *HTTPLog
	}{
		{"Build", HTTPLog{"ID", "www.netsparker.com", 0, 0, 0, 0, 0, 0, time.Now()}, args{target, refRequest, nil}, refHTTP},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.h.Build(tt.args.target, tt.args.request, tt.args.response)

			if got.TargetID != tt.want.TargetID || got.RequestURI != tt.want.RequestURI {
				t.Errorf("HTTPLog.Build() = %v, want %v", got, tt.want)
			}
		})
	}
}
