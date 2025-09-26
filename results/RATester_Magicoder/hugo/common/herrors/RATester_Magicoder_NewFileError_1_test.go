package herrors

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewFileError_1(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want FileError
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

			if got := NewFileError(tt.args.err); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFileError() = %v, want %v", got, tt.want)
			}
		})
	}
}
