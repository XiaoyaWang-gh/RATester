package herrors

import (
	"fmt"
	"reflect"
	"testing"
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
			name: "nil error",
			args: args{
				err: nil,
			},
			want: nil,
		},
		{
			name: "single error",
			args: args{
				err: fmt.Errorf("single error: %w", &fileError{}),
			},
			want: []FileError{&fileError{}},
		},
		{
			name: "multiple errors",
			args: args{
				err: fmt.Errorf("error 1: %w", &fileError{}),
			},
			want: []FileError{&fileError{}},
		},
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
