package validation

import (
	"errors"
	"fmt"
	"testing"
)

func TestAppendError_1(t *testing.T) {
	type args struct {
		err  error
		errs []error
	}
	tests := []struct {
		name string
		args args
		want error
	}{
		{
			name: "Test case 1",
			args: args{
				err:  errors.New("error 1"),
				errs: []error{errors.New("error 2"), errors.New("error 3")},
			},
			want: errors.Join(errors.New("error 1"), errors.New("error 2"), errors.New("error 3")),
		},
		{
			name: "Test case 2",
			args: args{
				err:  nil,
				errs: []error{errors.New("error 1"), errors.New("error 2")},
			},
			want: errors.Join(errors.New("error 1"), errors.New("error 2")),
		},
		{
			name: "Test case 3",
			args: args{
				err:  errors.New("error 1"),
				errs: []error{},
			},
			want: errors.New("error 1"),
		},
		{
			name: "Test case 4",
			args: args{
				err:  nil,
				errs: []error{},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := AppendError(tt.args.err, tt.args.errs...); !errors.Is(got, tt.want) {
				t.Errorf("AppendError() = %v, want %v", got, tt.want)
			}
		})
	}
}
