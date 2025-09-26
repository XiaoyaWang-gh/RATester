package schema

import (
	"fmt"
	"reflect"
	"testing"
)

func TestConvertFloat32_1(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name string
		args args
		want reflect.Value
	}{
		{
			name: "0",
			args: args{value: "1.1"},
			want: reflect.ValueOf(float32(1.1)),
		},
		{
			name: "1",
			args: args{value: "1.1"},
			want: reflect.ValueOf(float32(1.1)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := convertFloat32(tt.args.value); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertFloat32() = %v, want %v", got, tt.want)
			}
		})
	}
}
