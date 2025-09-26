package herrors

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewFileErrorFromName_1(t *testing.T) {
	type args struct {
		err  error
		name string
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

			if got := NewFileErrorFromName(tt.args.err, tt.args.name); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFileErrorFromName() = %v, want %v", got, tt.want)
			}
		})
	}
}
