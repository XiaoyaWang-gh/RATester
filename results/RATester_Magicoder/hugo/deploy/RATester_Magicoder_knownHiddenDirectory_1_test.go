package deploy

import (
	"fmt"
	"testing"
)

func TestknownHiddenDirectory_1(t *testing.T) {
	tests := []struct {
		name string
		want bool
	}{
		{".well-known", true},
		{"not-well-known", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := knownHiddenDirectory(tt.name); got != tt.want {
				t.Errorf("knownHiddenDirectory() = %v, want %v", got, tt.want)
			}
		})
	}
}
