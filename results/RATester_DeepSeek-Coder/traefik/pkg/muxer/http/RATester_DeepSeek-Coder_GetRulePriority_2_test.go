package http

import (
	"fmt"
	"testing"
)

func TestGetRulePriority_2(t *testing.T) {
	tests := []struct {
		name string
		rule string
		want int
	}{
		{"Test1", "*", 1},
		{"Test2", "**", 2},
		{"Test3", "***", 3},
		{"Test4", "****", 4},
		{"Test5", "*****", 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := GetRulePriority(tt.rule); got != tt.want {
				t.Errorf("GetRulePriority() = %v, want %v", got, tt.want)
			}
		})
	}
}
