package config

import (
	"fmt"
	"testing"
)

func TestString_3(t *testing.T) {
	tests := []struct {
		name string
		f    *BoolFuncFlag
		want string
	}{
		{
			name: "Test case 1",
			f: &BoolFuncFlag{
				v: true,
			},
			want: "true",
		},
		{
			name: "Test case 2",
			f: &BoolFuncFlag{
				v: false,
			},
			want: "false",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.f.String(); got != tt.want {
				t.Errorf("String() = %v, want %v", got, tt.want)
			}
		})
	}
}
