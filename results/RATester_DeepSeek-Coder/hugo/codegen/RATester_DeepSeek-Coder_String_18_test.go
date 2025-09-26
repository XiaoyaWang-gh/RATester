package codegen

import (
	"fmt"
	"reflect"
	"testing"
)

func TestString_18(t *testing.T) {
	type fields struct {
		Owner     reflect.Type
		OwnerName string
		Name      string
		Imports   []string
		In        []string
		Out       []string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "TestString",
			fields: fields{
				Owner:     reflect.TypeOf(Method{}),
				OwnerName: "Method",
				Name:      "String",
				Imports:   []string{},
				In:        []string{},
				Out:       []string{"string"},
			},
			want: "String() string {\n\treturn m.Name + m.inStr() + \" \" + m.outStr() + \"\\n\"\n}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			m := Method{
				Owner:     tt.fields.Owner,
				OwnerName: tt.fields.OwnerName,
				Name:      tt.fields.Name,
				Imports:   tt.fields.Imports,
				In:        tt.fields.In,
				Out:       tt.fields.Out,
			}
			if got := m.String(); got != tt.want {
				t.Errorf("Method.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
