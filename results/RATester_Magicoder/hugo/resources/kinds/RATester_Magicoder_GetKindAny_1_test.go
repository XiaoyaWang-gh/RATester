package kinds

import (
	"fmt"
	"testing"
)

func TestGetKindAny_1(t *testing.T) {
	tests := []struct {
		name string
		args string
		want string
	}{
		{"Test Case 1", "test1", "expected result 1"},
		{"Test Case 2", "test2", "expected result 2"},
		// Add more test cases as needed
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := GetKindAny(tt.args); got != tt.want {
				t.Errorf("GetKindAny() = %v, want %v", got, tt.want)
			}
		})
	}
}
