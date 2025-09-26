package template

import (
	"fmt"
	"sync"
	"testing"

	"github.com/gohugoio/hugo/tpl/internal/go_templates/texttemplate/parse"
)

func TestAssociate_1(t *testing.T) {
	type args struct {
		new  *Template
		tree *parse.Tree
	}
	tests := []struct {
		name string
		t    *Template
		args args
		want bool
	}{
		{
			name: "case1",
			t: &Template{
				name: "t",
				Tree: &parse.Tree{
					Name: "t",
				},
				common: &common{
					muTmpl: sync.RWMutex{},
					tmpl:   map[string]*Template{},
				},
			},
			args: args{
				new: &Template{
					name: "t",
					Tree: &parse.Tree{
						Name: "t",
					},
					common: &common{
						muTmpl: sync.RWMutex{},
						tmpl:   map[string]*Template{},
					},
				},
				tree: &parse.Tree{
					Name: "t",
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.t.associate(tt.args.new, tt.args.tree); got != tt.want {
				t.Errorf("associate() = %v, want %v", got, tt.want)
			}
		})
	}
}
