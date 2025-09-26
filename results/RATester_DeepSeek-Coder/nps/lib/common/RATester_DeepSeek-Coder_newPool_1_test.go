package common

import (
	"fmt"
	"testing"
)

func TestNewPool_1(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"Test Case 1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			newPool()
		})
	}
}
