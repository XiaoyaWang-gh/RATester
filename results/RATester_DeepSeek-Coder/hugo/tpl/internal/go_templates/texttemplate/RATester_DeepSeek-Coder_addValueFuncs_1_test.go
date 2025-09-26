package template

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAddValueFuncs_1(t *testing.T) {
	type testCase struct {
		name    string
		out     map[string]reflect.Value
		in      FuncMap
		wantErr bool
	}

	tests := []testCase{
		{
			name: "valid function",
			out:  make(map[string]reflect.Value),
			in: FuncMap{
				"validFunc": func() {},
			},
			wantErr: false,
		},
		{
			name: "invalid function",
			out:  make(map[string]reflect.Value),
			in: FuncMap{
				"invalidFunc": "not a function",
			},
			wantErr: true,
		},
		{
			name: "non-identifier function name",
			out:  make(map[string]reflect.Value),
			in: FuncMap{
				"1nvalidFunc": func() {},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			defer func() {
				if r := recover(); (r != nil) != tt.wantErr {
					t.Errorf("addValueFuncs() = %v, wantErr %v", r, tt.wantErr)
				}
			}()
			addValueFuncs(tt.out, tt.in)
		})
	}
}
