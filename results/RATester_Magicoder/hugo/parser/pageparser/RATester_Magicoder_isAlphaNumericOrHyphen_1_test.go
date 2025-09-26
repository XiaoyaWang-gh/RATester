package pageparser

import (
	"fmt"
	"testing"
)

func TestisAlphaNumericOrHyphen_1(t *testing.T) {

	tests := []struct {
		name string
		r    rune
		want bool
	}{
		{"Test1", 'a', true},
		{"Test2", '1', true},
		{"Test3", '-', true},
		{"Test4", '@', false},
		{"Test5", ' ', false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := isAlphaNumericOrHyphen(tt.r); got != tt.want {
				t.Errorf("isAlphaNumericOrHyphen() = %v, want %v", got, tt.want)
			}
		})
	}
}
