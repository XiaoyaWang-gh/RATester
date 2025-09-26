package template

import (
	"fmt"
	"reflect"
	"testing"
)

func TestBuiltinFuncs_1(t *testing.T) {
	tests := []struct {
		name string
		want map[string]reflect.Value
	}{
		{
			name: "Test builtinFuncs",
			want: createValueFuncs(builtins()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := builtinFuncs(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("builtinFuncs() = %v, want %v", got, tt.want)
			}
		})
	}
}
