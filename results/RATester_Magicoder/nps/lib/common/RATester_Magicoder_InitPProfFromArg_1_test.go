package common

import (
	"fmt"
	"testing"
)

func TestInitPProfFromArg_1(t *testing.T) {
	tests := []struct {
		name string
		arg  string
	}{
		{"Test case 1", "test1"},
		{"Test case 2", "test2"},
		{"Test case 3", ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			InitPProfFromArg(tt.arg)
		})
	}
}
