package herrors

import (
	"errors"
	"fmt"
	"os"
	"testing"
)

func TestIsNotExist_1(t *testing.T) {
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
				err: fmt.Errorf("some error: %w", os.ErrNotExist),
			},
			want: true,
		},
		{
			name: "Test case 3",
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

			if got := IsNotExist(tt.args.err); got != tt.want {
				t.Errorf("IsNotExist() = %v, want %v", got, tt.want)
			}
		})
	}
}
