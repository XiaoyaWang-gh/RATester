package modules

import (
	"fmt"
	"testing"
)

func TestVendor_1(t *testing.T) {
	tests := []struct {
		name string
		m    *moduleAdapter
		want bool
	}{
		{
			name: "Test case 1",
			m: &moduleAdapter{
				vendor: true,
			},
			want: true,
		},
		{
			name: "Test case 2",
			m: &moduleAdapter{
				vendor: false,
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.m.Vendor(); got != tt.want {
				t.Errorf("moduleAdapter.Vendor() = %v, want %v", got, tt.want)
			}
		})
	}
}
