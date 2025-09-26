package util

import (
	"fmt"
	"testing"
)

func TestRandIDWithLen_1(t *testing.T) {
	tests := []struct {
		name    string
		idLen   int
		wantErr bool
	}{
		{"Test case 1", 10, false},
		{"Test case 2", 0, false},
		{"Test case 3", -1, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			got, err := RandIDWithLen(tt.idLen)
			if (err != nil) != tt.wantErr {
				t.Errorf("RandIDWithLen() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) != tt.idLen {
				t.Errorf("RandIDWithLen() = %v, want %v", got, tt.idLen)
			}
		})
	}
}
