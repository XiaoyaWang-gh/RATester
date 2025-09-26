package herrors

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/spf13/afero"
)

func TestNewFileErrorFromFile_1(t *testing.T) {
	type args struct {
		err         error
		filename    string
		fs          afero.Fs
		linematcher LineMatcherFn
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

			if got := NewFileErrorFromFile(tt.args.err, tt.args.filename, tt.args.fs, tt.args.linematcher); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFileErrorFromFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
