package web

import (
	"fmt"
	"testing"
)

func TestGetPattern_1(t *testing.T) {
	tests := []struct {
		name string
		c    *ControllerInfo
		want string
	}{
		{
			name: "Test Case 1",
			c: &ControllerInfo{
				pattern: "test_pattern",
			},
			want: "test_pattern",
		},
		{
			name: "Test Case 2",
			c: &ControllerInfo{
				pattern: "another_pattern",
			},
			want: "another_pattern",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.c.GetPattern(); got != tt.want {
				t.Errorf("ControllerInfo.GetPattern() = %v, want %v", got, tt.want)
			}
		})
	}
}
