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
		{
			name: "Test case 1",
			rule: "rule1",
			want: 4,
		},
		{
			name: "Test case 2",
			rule: "rule2",
			want: 5,
		},
		// Add more test cases as needed
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
