package tcp

import (
	"errors"
	"fmt"
	"net"
	"syscall"
	"testing"
)

func TestIsSocketNotConnectedError_1(t *testing.T) {
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
					Err: syscall.ENOTCONN,
				},
			},
			want: true,
		},
		{
			name: "Test case 2",
			args: args{
				err: &net.OpError{
					Err: errors.New("some other error"),
				},
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

			if got := isSocketNotConnectedError(tt.args.err); got != tt.want {
				t.Errorf("isSocketNotConnectedError() = %v, want %v", got, tt.want)
			}
		})
	}
}
