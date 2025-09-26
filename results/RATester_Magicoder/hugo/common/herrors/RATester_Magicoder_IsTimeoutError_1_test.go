package herrors

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

func TestIsTimeoutError_1(t *testing.T) {
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
				err: &TimeoutError{
					Duration: 10 * time.Second,
				},
			},
			want: true,
		},
		{
			name: "Test case 2",
			args: args{
				err: errors.New("some error"),
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

			if got := IsTimeoutError(tt.args.err); got != tt.want {
				t.Errorf("IsTimeoutError() = %v, want %v", got, tt.want)
			}
		})
	}
}
