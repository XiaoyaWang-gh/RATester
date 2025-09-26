package filenotify

import (
	"fmt"
	"os"
	"testing"

	"github.com/fsnotify/fsnotify"
)

func TestCheckChange_1(t *testing.T) {
	type args struct {
		fi1 os.FileInfo
		fi2 os.FileInfo
	}
	tests := []struct {
		name string
		args args
		want fsnotify.Op
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := checkChange(tt.args.fi1, tt.args.fi2); got != tt.want {
				t.Errorf("checkChange() = %v, want %v", got, tt.want)
			}
		})
	}
}
