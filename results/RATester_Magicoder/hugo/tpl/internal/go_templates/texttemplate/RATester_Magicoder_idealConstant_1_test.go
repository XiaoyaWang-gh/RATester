package template

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/tpl/internal/go_templates/texttemplate/parse"
)

func TestidealConstant_1(t *testing.T) {
	tests := []struct {
		name     string
		constant *parse.NumberNode
		want     reflect.Value
	}{
		{
			name: "Test idealConstant with complex number",
			constant: &parse.NumberNode{
				IsComplex:  true,
				Complex128: complex(1, 2),
			},
			want: reflect.ValueOf(complex(1, 2)),
		},
		{
			name: "Test idealConstant with float number",
			constant: &parse.NumberNode{
				IsFloat: true,
				Float64: 1.23,
			},
			want: reflect.ValueOf(1.23),
		},
		{
			name: "Test idealConstant with int number",
			constant: &parse.NumberNode{
				IsInt: true,
				Int64: 123,
			},
			want: reflect.ValueOf(123),
		},
		{
			name: "Test idealConstant with uint number",
			constant: &parse.NumberNode{
				IsUint: true,
				Uint64: 123,
			},
			want: reflect.ValueOf(123),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			s := &state{}
			got := s.idealConstant(tt.constant)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("idealConstant() = %v, want %v", got, tt.want)
			}
		})
	}
}
