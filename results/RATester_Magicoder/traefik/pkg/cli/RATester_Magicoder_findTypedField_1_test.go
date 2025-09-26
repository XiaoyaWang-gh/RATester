package cli

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/traefik/paerser/parser"
)

func TestFindTypedField_1(t *testing.T) {
	type args struct {
		rType reflect.Type
		node  *parser.Node
	}
	tests := []struct {
		name  string
		args  args
		want  reflect.StructField
		want1 bool
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

			got, got1 := findTypedField(tt.args.rType, tt.args.node)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findTypedField() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("findTypedField() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
