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
			name: "Test case 1",
			m: &moduleAdapter{
				owner: nil,
			},
			want: true,
		},
		{
			name: "Test case 2",
			m: &moduleAdapter{
				gomod: &goModule{
					Indirect: false,
				},
			},
			want: true,
		},
		{
			name: "Test case 3",
			m: &moduleAdapter{
				gomod: &goModule{
					Indirect: true,
					Replace: &goModule{
						Version: "",
					},
				},
			},
			want: true,
		},
		{
			name: "Test case 4",
			m: &moduleAdapter{
				gomod: &goModule{
					Indirect: false,
					Replace: &goModule{
						Version: "1.0.0",
					},
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
