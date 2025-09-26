package middleware

import (
	"fmt"
	"io"
	"testing"
)

func TestReset_1(t *testing.T) {
	type args struct {
		reader io.ReadCloser
	}
	tests := []struct {
		name string
		r    *limitedReader
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

			tt.r.Reset(tt.args.reader)
		})
	}
}
