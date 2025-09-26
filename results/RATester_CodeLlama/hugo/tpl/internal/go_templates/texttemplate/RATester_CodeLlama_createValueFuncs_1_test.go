package template

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCreateValueFuncs_1(t *testing.T) {
	type args struct {
		funcMap FuncMap
	}
	tests := []struct {
		name string
		args args
		want map[string]reflect.Value
	}{
		{
			name: "test_createValueFuncs_0",
			args: args{
				funcMap: FuncMap{
					"name": func() string {
						return "name"
					},
				},
			},
			want: map[string]reflect.Value{
				"name": reflect.ValueOf(func() string {
					return "name"
				}),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := createValueFuncs(tt.args.funcMap); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("createValueFuncs() = %v, want %v", got, tt.want)
			}
		})
	}
}
