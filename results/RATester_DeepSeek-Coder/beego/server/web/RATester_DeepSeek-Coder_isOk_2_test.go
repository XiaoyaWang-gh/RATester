package web

import (
	"fmt"
	"os"
	"testing"
)

func TestIsOk_2(t *testing.T) {
	type args struct {
		s  *serveContentHolder
		fi os.FileInfo
	}
	tests := []struct {
		name string
		args args
		want bool
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

			if got := isOk(tt.args.s, tt.args.fi); got != tt.want {
				t.Errorf("isOk() = %v, want %v", got, tt.want)
			}
		})
	}
}
