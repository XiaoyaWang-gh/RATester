package modules

import (
	"fmt"
	"testing"
)

func TestIsValid_2(t *testing.T) {
	tests := []struct {
		name string
		v    HugoVersion
		want bool
	}{
		{
			name: "valid version",
			v: HugoVersion{
				Min: "0.100.0",
				Max: "0.200.0",
			},
			want: true,
		},
		{
			name: "invalid version",
			v: HugoVersion{
				Min: "0.300.0",
				Max: "0.200.0",
			},
			want: false,
		},
		{
			name: "extended version",
			v: HugoVersion{
				Min:      "0.100.0",
				Max:      "0.200.0",
				Extended: true,
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

			if got := tt.v.IsValid(); got != tt.want {
				t.Errorf("HugoVersion.IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
