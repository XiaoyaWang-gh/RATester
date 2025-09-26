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
			name: "empty",
			v:    HugoVersion{},
			want: "",
		},
		{
			name: "min",
			v: HugoVersion{
				Min: "0.50.0",
			},
			want: "Min 0.50.0",
		},
		{
			name: "max",
			v: HugoVersion{
				Max: "0.50.0",
			},
			want: "Max 0.50.0",
		},
		{
			name: "min-max",
			v: HugoVersion{
				Min: "0.50.0",
				Max: "0.50.0",
			},
			want: "0.50.0/0.50.0",
		},
		{
			name: "extended",
			v: HugoVersion{
				Min:      "0.50.0",
				Max:      "0.50.0",
				Extended: true,
			},
			want: "0.50.0/0.50.0 extended",
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
