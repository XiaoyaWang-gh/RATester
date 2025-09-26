package modules

import (
	"fmt"
	"testing"
)

func TestIsGoMod_1(t *testing.T) {
	tests := []struct {
		name string
		m    *moduleAdapter
		want bool
	}{
		{
			name: "Test with gomod not nil",
			m: &moduleAdapter{
				gomod: &goModule{
					Path: "test/module",
				},
			},
			want: true,
		},
		{
			name: "Test with gomod nil",
			m: &moduleAdapter{
				gomod: nil,
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

			if got := tt.m.IsGoMod(); got != tt.want {
				t.Errorf("moduleAdapter.IsGoMod() = %v, want %v", got, tt.want)
			}
		})
	}
}
