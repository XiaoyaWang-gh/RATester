package tcp

import (
	"fmt"
	"io"
	"net"
	"syscall"
	"testing"
)

func TestIsReadConnResetError_1(t *testing.T) {
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
				err: &net.OpError{
					Op:   "read",
					Err:  syscall.ECONNRESET,
					Addr: nil,
				},
			},
			want: true,
		},
		{
			name: "Test case 2",
			args: args{
				err: &net.OpError{
					Op:   "write",
					Err:  syscall.ECONNRESET,
					Addr: nil,
				},
			},
			want: false,
		},
		{
			name: "Test case 3",
			args: args{
				err: &net.OpError{
					Op:   "read",
					Err:  io.EOF,
					Addr: nil,
				},
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

			if got := isReadConnResetError(tt.args.err); got != tt.want {
				t.Errorf("isReadConnResetError() = %v, want %v", got, tt.want)
			}
		})
	}
}
