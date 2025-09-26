package fiber

import (
	"fmt"
	"testing"
)

func TestIsMethodIdempotent_1(t *testing.T) {
	tests := []struct {
		name string
		args string
		want bool
	}{
		{"GET is idempotent", MethodGet, true},
		{"PUT is idempotent", "PUT", true},
		{"DELETE is idempotent", "DELETE", true},
		{"POST is not idempotent", "POST", false},
		{"HEAD is idempotent", "HEAD", true},
		{"OPTIONS is idempotent", "OPTIONS", true},
		{"PATCH is not idempotent", "PATCH", false},
		{"TRACE is idempotent", "TRACE", true},
		{"CONNECT is not idempotent", "CONNECT", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := IsMethodIdempotent(tt.args); got != tt.want {
				t.Errorf("IsMethodIdempotent() = %v, want %v", got, tt.want)
			}
		})
	}
}
