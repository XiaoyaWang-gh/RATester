package logs

import (
	"fmt"
	"testing"
)

func TestColorByMethod_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var tests = []struct {
		method string
		want   string
	}{
		{"GET", "\x1b[32m"},
		{"POST", "\x1b[31m"},
		{"PUT", "\x1b[33m"},
		{"DELETE", "\x1b[31m"},
		{"PATCH", "\x1b[33m"},
		{"HEAD", "\x1b[37m"},
		{"OPTIONS", "\x1b[34m"},
		{"CONNECT", "\x1b[35m"},
		{"TRACE", "\x1b[36m"},
		{"", "\x1b[0m"},
	}
	for _, tt := range tests {
		if got := ColorByMethod(tt.method); got != tt.want {
			t.Errorf("ColorByMethod(%q) = %q, want %q", tt.method, got, tt.want)
		}
	}
}
