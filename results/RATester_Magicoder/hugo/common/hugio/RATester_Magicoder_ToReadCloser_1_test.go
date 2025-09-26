package hugio

import (
	"fmt"
	"io"
	"reflect"
	"testing"
)

func TestToReadCloser_1(t *testing.T) {
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name string
		args args
		want io.ReadCloser
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

			if got := ToReadCloser(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ToReadCloser() = %v, want %v", got, tt.want)
			}
		})
	}
}
