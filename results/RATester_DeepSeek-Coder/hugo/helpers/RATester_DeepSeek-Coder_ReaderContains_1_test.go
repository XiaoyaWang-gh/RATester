package helpers

import (
	"fmt"
	"io"
	"testing"
)

func TestReaderContains_1(t *testing.T) {
	type args struct {
		r        io.Reader
		subslice []byte
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

			if got := ReaderContains(tt.args.r, tt.args.subslice); got != tt.want {
				t.Errorf("ReaderContains() = %v, want %v", got, tt.want)
			}
		})
	}
}
