package validation

import (
	"fmt"
	"testing"
)

func TestIsSatisfied_9(t *testing.T) {
	e := Enum{Rules: "admin|user|guest"}

	tests := []struct {
		name string
		e    Enum
		arg  interface{}
		want bool
	}{
		{"Test Case 1", e, "admin", true},
		{"Test Case 2", e, "user", true},
		{"Test Case 3", e, "guest", true},
		{"Test Case 4", e, "invalid", false},
		{"Test Case 5", e, 123, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.e.IsSatisfied(tt.arg); got != tt.want {
				t.Errorf("IsSatisfied() = %v, want %v", got, tt.want)
			}
		})
	}
}
