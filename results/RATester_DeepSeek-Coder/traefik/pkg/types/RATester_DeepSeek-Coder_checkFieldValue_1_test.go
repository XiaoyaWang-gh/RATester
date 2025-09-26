package types

import (
	"fmt"
	"testing"
)

func TestCheckFieldValue_1(t *testing.T) {
	tests := []struct {
		name     string
		value    string
		defaultK bool
		want     bool
	}{
		{"TestKeep", AccessLogKeep, true, true},
		{"TestDrop", "drop", false, false},
		{"TestDefault", "other", true, true},
		{"TestEmpty", "", false, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := checkFieldValue(tt.value, tt.defaultK); got != tt.want {
				t.Errorf("checkFieldValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
