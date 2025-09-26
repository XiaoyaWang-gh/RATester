package modules

import (
	"fmt"
	"testing"
)

func TestString_12(t *testing.T) {
	tests := []struct {
		name string
		v    HugoVersion
		want string
	}{
		{
			name: "Test with Min and Max",
			v: HugoVersion{
				Min: "0.100.0",
				Max: "0.102.0",
			},
			want: "0.100.0/0.102.0",
		},
		{
			name: "Test with Min only",
			v: HugoVersion{
				Min: "0.100.0",
			},
			want: "Min 0.100.0",
		},
		{
			name: "Test with Max only",
			v: HugoVersion{
				Max: "0.102.0",
			},
			want: "Max 0.102.0",
		},
		{
			name: "Test with Extended",
			v: HugoVersion{
				Extended: true,
			},
			want: " extended",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.v.String(); got != tt.want {
				t.Errorf("HugoVersion.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
