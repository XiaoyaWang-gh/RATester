package schema

import (
	"fmt"
	"reflect"
	"testing"
)

func TestConvertBool_1(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
		want reflect.Value
	}{
		{
			name: "on",
			args: args{value: "on"},
			want: reflect.ValueOf(true),
		},
		{
			name: "true",
			args: args{value: "true"},
			want: reflect.ValueOf(true),
		},
		{
			name: "false",
			args: args{value: "false"},
			want: reflect.ValueOf(false),
		},
		{
			name: "invalid",
			args: args{value: "invalid"},
			want: invalidValue,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := convertBool(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertBool() = %v, want %v", got, tt.want)
			}
		})
	}
}
