package kv

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetFieldNames_1(t *testing.T) {
	type args struct {
		rootName string
		rootType reflect.Type
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "case1",
			args: args{
				rootName: "rootName",
				rootType: reflect.TypeOf(struct {
					Field1 string
					Field2 int
					Field3 struct {
						Field31 string
						Field32 int
					}
				}{}),
			},
			want: []string{
				"rootName.Field1",
				"rootName.Field2",
				"rootName.Field3.Field31",
				"rootName.Field3.Field32",
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

			if got := getFieldNames(tt.args.rootName, tt.args.rootType); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getFieldNames() = %v, want %v", got, tt.want)
			}
		})
	}
}
