package schema

import (
	"fmt"
	"reflect"
	"testing"
)

func TestIndirectType_1(t *testing.T) {
	type args struct {
		typ reflect.Type
	}
	tests := []struct {
		name string
		args args
		want reflect.Type
	}{
		{
			name: "test",
			args: args{typ: reflect.TypeOf(1)},
			want: reflect.TypeOf(1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := indirectType(tt.args.typ); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("indirectType() = %v, want %v", got, tt.want)
			}
		})
	}
}
