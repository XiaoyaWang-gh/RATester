package common

import (
	"fmt"
	"testing"
)

func TestGetBoolByStr_1(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want bool
	}{
		{"Test case 1", "1", true},
		{"Test case 2", "true", true},
		{"Test case 3", "0", false},
		{"Test case 4", "false", false},
		{"Test case 5", "random", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := GetBoolByStr(tt.arg); got != tt.want {
				t.Errorf("GetBoolByStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
