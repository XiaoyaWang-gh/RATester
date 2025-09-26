package client

import (
	"fmt"
	"testing"
)

func TestStartFromFile_1(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{"TestStartFromFile", args{"/path/to/config/file"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			StartFromFile(tt.args.path)
		})
	}
}
