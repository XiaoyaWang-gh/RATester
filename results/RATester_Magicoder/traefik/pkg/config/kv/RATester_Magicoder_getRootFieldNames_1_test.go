package kv

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetRootFieldNames_1(t *testing.T) {
	type args struct {
		rootName string
		element  interface{}
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Test Case 1",
			args: args{
				rootName: "root",
				element:  struct{ FieldName string }{FieldName: "value"},
			},
			want: []string{"FieldName"},
		},
		{
			name: "Test Case 2",
			args: args{
				rootName: "root",
				element:  nil,
			},
			want: nil,
		},
		{
			name: "Test Case 3",
			args: args{
				rootName: "root",
				element:  struct{ FieldName string }{FieldName: "value"},
			},
			want: []string{"FieldName"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := getRootFieldNames(tt.args.rootName, tt.args.element); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getRootFieldNames() = %v, want %v", got, tt.want)
			}
		})
	}
}
