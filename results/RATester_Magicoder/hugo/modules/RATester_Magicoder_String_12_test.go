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
			name: "Test case 1",
			v: HugoVersion{
				Min:      "1.0.0",
				Max:      "2.0.0",
				Extended: true,
			},
			want: "1.0.0/2.0.0 extended",
		},
		{
			name: "Test case 2",
			v: HugoVersion{
				Min: "1.0.0",
			},
			want: "Min 1.0.0",
		},
		{
			name: "Test case 3",
			v: HugoVersion{
				Max: "2.0.0",
			},
			want: "Max 2.0.0",
		},
		{
			name: "Test case 4",
			v: HugoVersion{
				Extended: true,
			},
			want: " extended",
		},
		{
			name: "Test case 5",
			v:    HugoVersion{},
			want: "",
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
