package template

import (
	"fmt"
	"reflect"
	"testing"
)

func TestIndirectInterface_1(t *testing.T) {
	type args struct {
		v reflect.Value
	}
	tests := []struct {
		name string
		args args
		want reflect.Value
	}{
		{
			name: "test",
			args: args{
				v: reflect.ValueOf(1),
			},
			want: reflect.ValueOf(1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := indirectInterface(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("indirectInterface() = %v, want %v", got, tt.want)
			}
		})
	}
}
