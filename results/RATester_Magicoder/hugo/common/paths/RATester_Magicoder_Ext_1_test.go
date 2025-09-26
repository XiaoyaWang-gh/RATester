package paths

import (
	"fmt"
	"testing"
)

func TestExt_1(t *testing.T) {
	bridge := pathBridge{}
	tests := []struct {
		name string
		in   string
		want string
	}{
		{"empty", "", ""},
		{"root", "/", ""},
		{"no extension", "file", ""},
		{"with extension", "file.txt", ".txt"},
		{"with multiple dots", "file.with.dots.txt", ".txt"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := bridge.Ext(tt.in); got != tt.want {
				t.Errorf("Ext() = %v, want %v", got, tt.want)
			}
		})
	}
}
