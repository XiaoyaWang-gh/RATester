package client

import (
	"fmt"
	"io"
	"reflect"
	"testing"
)

func TestSetFileReader_1(t *testing.T) {
	type args struct {
		r io.ReadCloser
	}
	tests := []struct {
		name string
		args args
		want SetFileFunc
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

			if got := SetFileReader(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SetFileReader() = %v, want %v", got, tt.want)
			}
		})
	}
}
