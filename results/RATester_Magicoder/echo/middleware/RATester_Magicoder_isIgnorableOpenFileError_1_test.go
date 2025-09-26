package middleware

import (
	"fmt"
	"os"
	"testing"
)

func TestIsIgnorableOpenFileError_1(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test case 1",
			args: args{
				err: os.ErrNotExist,
			},
			want: true,
		},
		{
			name: "Test case 2",
			args: args{
				err: os.ErrPermission,
			},
			want: false,
		},
		{
			name: "Test case 3",
			args: args{
				err: nil,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := isIgnorableOpenFileError(tt.args.err); got != tt.want {
				t.Errorf("isIgnorableOpenFileError() = %v, want %v", got, tt.want)
			}
		})
	}
}
