package admin

import (
	"fmt"
	"io"
	"testing"
)

func TestGetCPUProfile_1(t *testing.T) {
	type args struct {
		w io.Writer
	}
	tests := []struct {
		name string
		args args
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

			GetCPUProfile(tt.args.w)
		})
	}
}
