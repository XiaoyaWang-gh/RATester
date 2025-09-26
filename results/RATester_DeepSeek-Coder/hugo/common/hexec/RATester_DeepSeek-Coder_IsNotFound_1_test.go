package hexec

import (
	"errors"
	"fmt"
	"testing"
)

func TestIsNotFound_1(t *testing.T) {
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
				err: errors.New("not found"),
			},
			want: false,
		},
		{
			name: "Test case 2",
			args: args{
				err: &NotFoundError{
					name:   "Test",
					method: "Get",
				},
			},
			want: true,
		},
		{
			name: "Test case 3",
			args: args{
				err: errors.New("not found"),
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

			if got := IsNotFound(tt.args.err); got != tt.want {
				t.Errorf("IsNotFound() = %v, want %v", got, tt.want)
			}
		})
	}
}
