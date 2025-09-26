package config

import (
	"fmt"
	"testing"
)

func TestStringToGibabyte_1(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want uint64
	}{
		{"Test case 1", "1", 1073741824},
		{"Test case 2", "0", 0},
		{"Test case 3", "-1", 0},
		{"Test case 4", "1.5", 1610612736},
		{"Test case 5", "invalid", 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := stringToGibabyte(tt.arg); got != tt.want {
				t.Errorf("stringToGibabyte() = %v, want %v", got, tt.want)
			}
		})
	}
}
