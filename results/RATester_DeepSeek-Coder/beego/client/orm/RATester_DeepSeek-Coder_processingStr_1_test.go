package orm

import (
	"fmt"
	"testing"
)

func TestProcessingStr_1(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want string
	}{
		{
			name: "Test Case 1",
			args: []string{"test1", "test2", "test3"},
			want: "\"test1\",\"test2\",\"test3\"",
		},
		{
			name: "Test Case 2",
			args: []string{"test4", "test5", "test6"},
			want: "\"test4\",\"test5\",\"test6\"",
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

			if got := processingStr(tt.args); got != tt.want {
				t.Errorf("processingStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
