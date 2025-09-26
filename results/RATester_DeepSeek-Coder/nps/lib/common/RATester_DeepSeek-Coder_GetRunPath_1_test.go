package common

import (
	"fmt"
	"testing"
)

func TestGetRunPath_1(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"Test1", "expected_path"},
		{"Test2", "expected_path"},
		{"Test3", "expected_path"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := GetRunPath(); got != tt.want {
				t.Errorf("GetRunPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
