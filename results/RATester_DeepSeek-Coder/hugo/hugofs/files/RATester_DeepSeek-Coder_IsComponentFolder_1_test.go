package files

import (
	"fmt"
	"testing"
)

func TestIsComponentFolder_1(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		{"test1", true},
		{"test2", false},
		{"test3", true},
		{"test4", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := IsComponentFolder(tt.name); got != tt.want {
				t.Errorf("IsComponentFolder() = %v, want %v", got, tt.want)
			}
		})
	}
}
