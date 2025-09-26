package testenv

import (
	"errors"
	"fmt"
	"io/fs"
	"syscall"
	"testing"
)

func TestsyscallIsNotSupported_1(t *testing.T) {
	tests := []struct {
		name string
		err  error
		want bool
	}{
		{
			name: "nil error",
			err:  nil,
			want: false,
		},
		{
			name: "syscall.EPERM",
			err:  syscall.EPERM,
			want: true,
		},
		{
			name: "syscall.EROFS",
			err:  syscall.EROFS,
			want: true,
		},
		{
			name: "syscall.EINVAL",
			err:  syscall.EINVAL,
			want: true,
		},
		{
			name: "fs.ErrPermission",
			err:  fs.ErrPermission,
			want: true,
		},
		{
			name: "errors.ErrUnsupported",
			err:  errors.ErrUnsupported,
			want: true,
		},
		{
			name: "random error",
			err:  errors.New("random error"),
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

			if got := syscallIsNotSupported(tt.err); got != tt.want {
				t.Errorf("syscallIsNotSupported() = %v, want %v", got, tt.want)
			}
		})
	}
}
