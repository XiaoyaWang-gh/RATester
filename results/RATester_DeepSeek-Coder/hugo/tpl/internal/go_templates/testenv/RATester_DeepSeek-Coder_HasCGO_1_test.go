package testenv

import (
	"fmt"
	"testing"
)

func TestHasCGO_1(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := HasCGO(); got != tt.want {
				t.Errorf("HasCGO() = %v, want %v", got, tt.want)
			}
		})
	}
}
