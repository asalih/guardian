package models

import (
	"net/http"
	"reflect"
	"testing"
	"time"
)

func TestNewHTTPLog(t *testing.T) {
	tests := []struct {
		name string
		want *HTTPLog
	}{
		{"Init", &HTTPLog{"", "", 0, 0, 0, 0, 0, 0, time.Now()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewHTTPLog(); !reflect.DeepEqual(got, tt.want) {
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
	refRequest, _ := http.NewRequest("GET", "www.netsparker.com", nil)
	target := new(Target)

	tests := []struct {
		name string
		h    HTTPLog
		args args
		want *HTTPLog
	}{
		{"Build", HTTPLog{"", "", 0, 0, 0, 0, 0, 0, time.Now()}, args{target, refRequest, nil}, refHTTP},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.h.Build(tt.args.target, tt.args.request, tt.args.response); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HTTPLog.Build() = %v, want %v", got, tt.want)
			}
		})
	}
}
