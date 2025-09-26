package testenv

import (
	"fmt"
	"testing"
)

func TestHasParallelism_1(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		{"Test1", true},
		{"Test2", true},
		{"Test3", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := HasParallelism(); got != tt.want {
				t.Errorf("HasParallelism() = %v, want %v", got, tt.want)
			}
		})
	}
}
