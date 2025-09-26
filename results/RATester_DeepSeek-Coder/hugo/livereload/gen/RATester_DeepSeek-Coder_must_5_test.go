package main

import (
	"errors"
	"fmt"
	"testing"
)

func TestMust_5(t *testing.T) {
	tests := []struct {
		name string
		err  error
	}{
		{
			name: "Test case 1",
			err:  nil,
		},
		{
			name: "Test case 2",
			err:  errors.New("some error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			defer func() {
				if r := recover(); r != nil {
					if tt.err != nil {
						t.Errorf("Expected no panic, but got: %v", r)
					}
				} else if tt.err == nil {
					t.Errorf("Expected panic, but got none")
				}
			}()

			must(tt.err)
		})
	}
}
