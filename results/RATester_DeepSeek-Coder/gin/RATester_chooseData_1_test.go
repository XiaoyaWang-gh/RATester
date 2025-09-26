package gin

import (
	"fmt"
	"testing"
)

func TestChooseData_1(t *testing.T) {
	type test struct {
		name     string
		custom   any
		wildcard any
		want     any
	}

	tests := []test{
		{
			name:     "both custom and wildcard are not nil",
			custom:   "custom",
			wildcard: "wildcard",
			want:     "custom",
		},
		{
			name:     "only custom is not nil",
			custom:   "custom",
			wildcard: nil,
			want:     "custom",
		},
		{
			name:     "only wildcard is not nil",
			custom:   nil,
			wildcard: "wildcard",
			want:     "wildcard",
		},
		{
			name:     "both custom and wildcard are nil",
			custom:   nil,
			wildcard: nil,
			want:     "panic: negotiation config is invalid",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			defer func() {
				if r := recover(); r != nil {
					if r != tt.want {
						t.Errorf("chooseData() = %v, want %v", r, tt.want)
					}
				}
			}()

			if got := chooseData(tt.custom, tt.wildcard); got != tt.want {
				t.Errorf("chooseData() = %v, want %v", got, tt.want)
			}
		})
	}
}
