package codegen

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDeclarationNamed_1(t *testing.T) {
	type fields struct {
		Owner     reflect.Type
		OwnerName string
		Name      string
		Imports   []string
		In        []string
		Out       []string
	}
	type args struct {
		receiver string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		m := Method{
			Owner:     tt.fields.Owner,
			OwnerName: tt.fields.OwnerName,
			Name:      tt.fields.Name,
			Imports:   tt.fields.Imports,
			In:        tt.fields.In,
			Out:       tt.fields.Out,
		}
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := m.DeclarationNamed(tt.args.receiver); got != tt.want {
				t.Errorf("Method.DeclarationNamed() = %v, want %v", got, tt.want)
			}
		})
	}
}
