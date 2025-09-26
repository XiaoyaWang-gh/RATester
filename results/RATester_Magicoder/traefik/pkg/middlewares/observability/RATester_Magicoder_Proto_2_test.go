package observability

import (
	"fmt"
	"testing"
)

func TestProto_2(t *testing.T) {
	tests := []struct {
		name  string
		proto string
		want  string
	}{
		{
			name:  "HTTP/1.0",
			proto: "HTTP/1.0",
			want:  "1.0",
		},
		{
			name:  "HTTP/1.1",
			proto: "HTTP/1.1",
			want:  "1.1",
		},
		{
			name:  "HTTP/2",
			proto: "HTTP/2",
			want:  "2",
		},
		{
			name:  "HTTP/3",
			proto: "HTTP/3",
			want:  "3",
		},
		{
			name:  "Unknown",
			proto: "Unknown",
			want:  "Unknown",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := Proto(tt.proto); got != tt.want {
				t.Errorf("Proto() = %v, want %v", got, tt.want)
			}
		})
	}
}
