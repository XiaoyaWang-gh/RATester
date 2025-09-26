package template

import (
	"fmt"
	"testing"
)

func TestGoodName_1(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		{"", false},
		{"test", true},
		{"test123", true},
		{"_test", false},
		{"test_", true},
		{"123test", false},
		{"test_123", true},
		{"Test", true},
		{"Test123", true},
		{"Test_123", true},
		{"Test_123_", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := goodName(tt.name); got != tt.want {
				t.Errorf("goodName() = %v, want %v", got, tt.want)
			}
		})
	}
}
