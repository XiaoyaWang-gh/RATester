package herrors

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/common/text"
)

func TestUnwrapFileErrors_1(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want []FileError
	}{
		{
			name: "Test case 1",
			args: args{
				err: errors.New("test error"),
			},
			want: []FileError{},
		},
		{
			name: "Test case 2",
			args: args{
				err: &fileError{
					position:     text.Position{},
					errorContext: &ErrorContext{},
					fileType:     "",
					cause:        errors.New("test error"),
				},
			},
			want: []FileError{
				&fileError{
					position:     text.Position{},
					errorContext: &ErrorContext{},
					fileType:     "",
					cause:        errors.New("test error"),
				},
			},
		},
		// Add more test cases as needed
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := UnwrapFileErrors(tt.args.err); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UnwrapFileErrors() = %v, want %v", got, tt.want)
			}
		})
	}
}
