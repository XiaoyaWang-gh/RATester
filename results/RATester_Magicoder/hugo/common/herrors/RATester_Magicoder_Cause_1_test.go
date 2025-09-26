package herrors

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func TestCause_1(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want error
	}{
		{
			name: "Test case 1",
			args: args{
				err: errors.New("test error"),
			},
			want: errors.New("test error"),
		},
		{
			name: "Test case 2",
			args: args{
				err: fmt.Errorf("wrapped: %w", errors.New("test error")),
			},
			want: errors.New("test error"),
		},
		{
			name: "Test case 3",
			args: args{
				err: nil,
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

			if got := Cause(tt.args.err); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Cause() = %v, want %v", got, tt.want)
			}
		})
	}
}
