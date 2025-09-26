package parse

import (
	"fmt"
	"strings"
	"testing"
)

func TestWriteTo_1(t *testing.T) {
	tests := []struct {
		name string
		t    *TemplateNode
		want string
	}{
		{
			name: "Test with nil Pipe",
			t: &TemplateNode{
				Name: "test",
				Pipe: nil,
			},
			want: "{{template \"test\"}}",
		},
		{
			name: "Test with non-nil Pipe",
			t: &TemplateNode{
				Name: "test",
				Pipe: &PipeNode{
					Decl: []*VariableNode{
						{
							Ident: []string{"$var1"},
						},
						{
							Ident: []string{"$var2", "field1"},
						},
					},
				},
			},
			want: "{{template \"test\" $var1 $var2.field1}}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			sb := &strings.Builder{}
			tt.t.writeTo(sb)
			if got := sb.String(); got != tt.want {
				t.Errorf("writeTo() = %v, want %v", got, tt.want)
			}
		})
	}
}
