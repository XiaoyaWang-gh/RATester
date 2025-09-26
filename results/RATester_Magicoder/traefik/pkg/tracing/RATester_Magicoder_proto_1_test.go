package tracing

import (
	"fmt"
	"testing"
)

func TestProto_1(t *testing.T) {
	tests := []struct {
		name  string
		proto string
		want  string
	}{
		{"HTTP/1.0", "HTTP/1.0", "1.0"},
		{"HTTP/1.1", "HTTP/1.1", "1.1"},
		{"HTTP/2", "HTTP/2", "2"},
		{"HTTP/3", "HTTP/3", "3"},
		{"Invalid", "Invalid", "Invalid"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := proto(tt.proto); got != tt.want {
				t.Errorf("proto() = %v, want %v", got, tt.want)
			}
		})
	}
}
