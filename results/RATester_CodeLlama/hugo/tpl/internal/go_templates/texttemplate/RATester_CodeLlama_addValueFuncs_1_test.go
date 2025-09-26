package template

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAddValueFuncs_1(t *testing.T) {
	type args struct {
		out map[string]reflect.Value
		in  FuncMap
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			addValueFuncs(tt.args.out, tt.args.in)
		})
	}
}
