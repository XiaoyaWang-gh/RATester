package utils

import (
	"fmt"
	"testing"
)

func TestisPrintable_1(t *testing.T) {
	tests := []struct {
		name string
		c    byte
		want bool
	}{
		{"Test1", 'A', true},
		{"Test2", 'a', true},
		{"Test3", '1', true},
		{"Test4", ' ', true},
		{"Test5", '\n', true},
		{"Test6", '\t', true},
		{"Test7", '', false},
		{"Test8", '', false},
		{"Test9", '', false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := isPrintable(tt.c); got != tt.want {
				t.Errorf("isPrintable() = %v, want %v", got, tt.want)
			}
		})
	}
}
