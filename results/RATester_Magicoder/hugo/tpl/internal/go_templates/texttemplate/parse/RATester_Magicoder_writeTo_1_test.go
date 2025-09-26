package parse

import (
	"fmt"
	"strings"
	"testing"
)

func TestwriteTo_1(t *testing.T) {
	tests := []struct {
		name string
		node *TemplateNode
		want string
	}{
		{
			name: "test1",
			node: &TemplateNode{
				Name: "test1",
			},
			want: "{{template \"test1\"}}",
		},
		{
			name: "test2",
			node: &TemplateNode{
				Name: "test2",
				Pipe: &PipeNode{
					Decl: []*VariableNode{
						{
							Ident: []string{"$var1"},
						},
					},
				},
			},
			want: "{{template \"test2\" $var1}}",
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
			tt.node.writeTo(sb)
			if got := sb.String(); got != tt.want {
				t.Errorf("writeTo() = %v, want %v", got, tt.want)
			}
		})
	}
}
