package commands

import (
	"errors"
	"fmt"
	"reflect"
	"testing"

	"github.com/bep/simplecobra"
)

func TestNewUserError_1(t *testing.T) {
	type args struct {
		a []any
	}
	tests := []struct {
		name string
		args args
		want *simplecobra.CommandError
	}{
		{
			name: "test",
			args: args{
				a: []any{
					"test",
				},
			},
			want: &simplecobra.CommandError{Err: errors.New("test")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := newUserError(tt.args.a...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newUserError() = %v, want %v", got, tt.want)
			}
		})
	}
}
