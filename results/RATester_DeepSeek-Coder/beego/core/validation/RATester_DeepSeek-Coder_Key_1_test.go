package validation

import (
	"fmt"
	"reflect"
	"testing"
)

func TestKey_1(t *testing.T) {
	type args struct {
		key string
	}
	tests := []struct {
		name string
		r    *Result
		args args
		want *Result
	}{
		{
			name: "TestKey_0",
			r:    &Result{Error: &Error{}, Ok: true},
			args: args{"key_0"},
			want: &Result{Error: &Error{Key: "key_0"}, Ok: true},
		},
		{
			name: "TestKey_1",
			r:    &Result{Error: nil, Ok: true},
			args: args{"key_1"},
			want: &Result{Error: nil, Ok: true},
		},
		{
			name: "TestKey_2",
			r:    &Result{Error: &Error{}, Ok: false},
			args: args{"key_2"},
			want: &Result{Error: &Error{}, Ok: false},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.r.Key(tt.args.key); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Result.Key() = %v, want %v", got, tt.want)
			}
		})
	}
}
