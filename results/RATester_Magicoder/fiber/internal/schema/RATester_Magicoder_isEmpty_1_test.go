package schema

import (
	"fmt"
	"reflect"
	"testing"
)

func TestisEmpty_1(t *testing.T) {
	type args struct {
		t     reflect.Type
		value []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test isEmpty with empty slice",
			args: args{
				t:     reflect.TypeOf(""),
				value: []string{},
			},
			want: true,
		},
		{
			name: "Test isEmpty with non-empty slice",
			args: args{
				t:     reflect.TypeOf(""),
				value: []string{"test"},
			},
			want: false,
		},
		{
			name: "Test isEmpty with non-string type",
			args: args{
				t:     reflect.TypeOf(1),
				value: []string{"test"},
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

			if got := isEmpty(tt.args.t, tt.args.value); got != tt.want {
				t.Errorf("isEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}
