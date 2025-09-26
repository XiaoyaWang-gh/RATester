package types

import (
	"fmt"
	"testing"
)

func TestBytes_1(t *testing.T) {
	tests := []struct {
		name string
		q    *BandwidthQuantity
		want int64
	}{
		{
			name: "MB",
			q:    &BandwidthQuantity{s: "MB", i: 1024 * 1024},
			want: 1024 * 1024,
		},
		{
			name: "KB",
			q:    &BandwidthQuantity{s: "KB", i: 1024},
			want: 1024,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.q.Bytes(); got != tt.want {
				t.Errorf("Bytes() = %v, want %v", got, tt.want)
			}
		})
	}
}
