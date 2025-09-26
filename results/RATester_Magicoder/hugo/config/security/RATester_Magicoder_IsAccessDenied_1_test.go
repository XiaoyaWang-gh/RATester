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
			name: "AccessDeniedError",
			args: args{
				err: &AccessDeniedError{
					name: "test",
				},
			},
			want: true,
		},
		{
			name: "OtherError",
			args: args{
				err: errors.New("other error"),
			},
			want: false,
		},
		{
			name: "NilError",
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

			if got := IsAccessDenied(tt.args.err); got != tt.want {
				t.Errorf("IsAccessDenied() = %v, want %v", got, tt.want)
			}
		})
	}
}
