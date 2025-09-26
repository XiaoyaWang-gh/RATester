package herrors

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/common/text"
)

func TestUnwrapFileError_1(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want FileError
	}{
		{
			name: "Test case 1",
			args: args{
				err: errors.New("test error"),
			},
			want: nil,
		},
		{
			name: "Test case 2",
			args: args{
				err: &fileError{
					position:     text.Position{},
					errorContext: &ErrorContext{},
					fileType:     "",
					cause:        nil,
				},
			},
			want: &fileError{
				position:     text.Position{},
				errorContext: &ErrorContext{},
				fileType:     "",
				cause:        nil,
			},
		},
		// Add more test cases here...
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := UnwrapFileError(tt.args.err); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UnwrapFileError() = %v, want %v", got, tt.want)
			}
		})
	}
}
