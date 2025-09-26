package hugio

import (
	"fmt"
	"io"
	"reflect"
	"testing"
)

func TestNewMultiWriteCloser_1(t *testing.T) {
	type args struct {
		writeClosers []io.WriteCloser
	}
	tests := []struct {
		name string
		args args
		want io.WriteCloser
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

			if got := NewMultiWriteCloser(tt.args.writeClosers...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMultiWriteCloser() = %v, want %v", got, tt.want)
			}
		})
	}
}
