package parse

import (
	"fmt"
	"testing"
)

func TestCheckPipeline_1(t *testing.T) {
	type args struct {
		pipe     *PipeNode
		context  string
		expected error
	}

	tree := &Tree{}

	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test case 1",
			args: args{
				pipe: &PipeNode{
					Cmds: []*CommandNode{
						{
							Args: []Node{
								&VariableNode{
									Ident: []string{"var1"},
								},
							},
						},
					},
				},
				context:  "context1",
				expected: nil,
			},
		},
		{
			name: "Test case 2",
			args: args{
				pipe: &PipeNode{
					Cmds: []*CommandNode{
						{
							Args: []Node{
								&VariableNode{
									Ident: []string{"var2"},
								},
							},
						},
					},
				},
				context:  "context2",
				expected: nil,
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

			tree.checkPipeline(tt.args.pipe, tt.args.context)
			// Add your assertions here
		})
	}
}
