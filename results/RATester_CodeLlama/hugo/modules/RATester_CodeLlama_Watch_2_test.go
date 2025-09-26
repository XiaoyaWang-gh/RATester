package modules

import (
	"fmt"
	"testing"
)

func TestWatch_2(t *testing.T) {
	tests := []struct {
		name string
		m    *moduleAdapter
		want bool
	}{
		{
			name: "test case 1",
			m: &moduleAdapter{
				gomod: &goModule{
					Indirect: true,
				},
			},
			want: true,
		},
		{
			name: "test case 2",
			m: &moduleAdapter{
				gomod: &goModule{
					Indirect: false,
				},
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

			if got := tt.m.Watch(); got != tt.want {
				t.Errorf("moduleAdapter.Watch() = %v, want %v", got, tt.want)
			}
		})
	}
}
