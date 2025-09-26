package web

import (
	"fmt"
	"reflect"
	"testing"
)

func Testlist_1(t *testing.T) {
	type args struct {
		root string
		p    interface{}
		m    M
	}
	tests := []struct {
		name string
		args args
		want M
	}{
		{
			name: "Test case 1",
			args: args{
				root: "",
				p: struct {
					FieldName string
				}{
					FieldName: "test",
				},
				m: make(M),
			},
			want: M{
				"FieldName": "test",
			},
		},
		{
			name: "Test case 2",
			args: args{
				root: "root",
				p: struct {
					FieldName string
				}{
					FieldName: "test",
				},
				m: make(M),
			},
			want: M{
				"root.FieldName": "test",
			},
		},
		{
			name: "Test case 3",
			args: args{
				root: "",
				p: struct {
					FieldName string
					SubField  struct {
						SubFieldName string
					}
				}{
					FieldName: "test",
					SubField: struct {
						SubFieldName string
					}{
						SubFieldName: "subtest",
					},
				},
				m: make(M),
			},
			want: M{
				"FieldName":             "test",
				"SubField.SubFieldName": "subtest",
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

			list(tt.args.root, tt.args.p, tt.args.m)
			if !reflect.DeepEqual(tt.args.m, tt.want) {
				t.Errorf("list() = %v, want %v", tt.args.m, tt.want)
			}
		})
	}
}
