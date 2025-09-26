package pandoc

import (
	"fmt"
	"testing"
)

func TestGetPandocBinaryName_1(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"Test Case 1", pandocBinary},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := getPandocBinaryName(); got != tt.want {
				t.Errorf("getPandocBinaryName() = %v, want %v", got, tt.want)
			}
		})
	}
}
