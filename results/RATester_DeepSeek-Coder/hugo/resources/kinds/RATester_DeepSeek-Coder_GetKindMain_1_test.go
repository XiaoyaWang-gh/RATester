package kinds

import (
	"fmt"
	"testing"
)

func TestGetKindMain_1(t *testing.T) {
	tests := []struct {
		name string
		args string
		want string
	}{
		{"Test1", "test1", "test1"},
		{"Test2", "test2", "test2"},
		{"Test3", "test3", "test3"},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := GetKindMain(tt.args); got != tt.want {
				t.Errorf("GetKindMain() = %v, want %v", got, tt.want)
			}
		})
	}
}
