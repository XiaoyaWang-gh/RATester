package client

import (
	"fmt"
	"io"
	"reflect"
	"testing"
)

func TestSetReader_1(t *testing.T) {
	type args struct {
		r io.ReadCloser
	}
	tests := []struct {
		name string
		f    *File
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

			tt.f.SetReader(tt.args.r)
			if !reflect.DeepEqual(tt.f.reader, tt.want) {
				t.Errorf("File.SetReader() = %v, want %v", tt.f.reader, tt.want)
			}
		})
	}
}
