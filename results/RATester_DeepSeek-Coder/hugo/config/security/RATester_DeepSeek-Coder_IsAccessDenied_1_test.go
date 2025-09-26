package security

import (
	"errors"
	"fmt"
	"testing"
)

func TestIsAccessDenied_1(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Access Denied",
			args: args{
				err: &AccessDeniedError{
					name: "test",
				},
			},
			want: true,
		},
		{
			name: "Access Not Denied",
			args: args{
				err: errors.New("test"),
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

			if got := IsAccessDenied(tt.args.err); got != tt.want {
				t.Errorf("IsAccessDenied() = %v, want %v", got, tt.want)
			}
		})
	}
}
