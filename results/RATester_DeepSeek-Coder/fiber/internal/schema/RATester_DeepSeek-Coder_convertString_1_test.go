package schema

import (
	"fmt"
	"reflect"
	"testing"
)

func TestConvertString_1(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
		want reflect.Value
	}{
		{
			name: "Test Case 1",
			args: args{
				value: "test",
			},
			want: reflect.ValueOf("test"),
		},
		{
			name: "Test Case 2",
			args: args{
				value: "test2",
			},
			want: reflect.ValueOf("test2"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := convertString(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertString() = %v, want %v", got, tt.want)
			}
		})
	}
}
