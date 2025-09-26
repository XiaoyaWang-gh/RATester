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
			name: "Test case 1",
			q:    &BandwidthQuantity{s: "MB", i: 1024},
			want: 1024,
		},
		{
			name: "Test case 2",
			q:    &BandwidthQuantity{s: "KB", i: 1024},
			want: 1024,
		},
		{
			name: "Test case 3",
			q:    &BandwidthQuantity{s: "MB", i: 2048},
			want: 2048,
		},
		{
			name: "Test case 4",
			q:    &BandwidthQuantity{s: "KB", i: 2048},
			want: 2048,
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
