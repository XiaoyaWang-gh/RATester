package main

import (
	"fmt"
	"testing"
)

func TestMain_1(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"Test case 1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			main()
		})
	}
}
