package paths

import (
	"fmt"
	"testing"
)

func TestHasExt_1(t *testing.T) {
	tests := []struct {
		name string
		arg  string
		want bool
	}{
		{"NoExt", "test", false},
		{"WithExt", "test.txt", true},
		{"WithExtInDir", "dir/test.txt", true},
		{"NoExtInDir", "dir/test", false},
		{"NoExtInDirWithSlash", "dir/test/", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := HasExt(tt.arg); got != tt.want {
				t.Errorf("HasExt() = %v, want %v", got, tt.want)
			}
		})
	}
}
