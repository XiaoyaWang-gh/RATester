package tcp

import (
	"fmt"
	"testing"
)

func TestVarGroupName_1(t *testing.T) {
	tests := []struct {
		name string
		idx  int
		want string
	}{
		{"Test case 1", 1, "v1"},
		{"Test case 2", 2, "v2"},
		{"Test case 3", 3, "v3"},
		// Add more test cases as needed
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := varGroupName(tt.idx); got != tt.want {
				t.Errorf("varGroupName() = %v, want %v", got, tt.want)
			}
		})
	}
}
