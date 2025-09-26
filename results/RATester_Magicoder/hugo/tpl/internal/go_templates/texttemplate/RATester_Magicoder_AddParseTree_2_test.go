package template

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/tpl/internal/go_templates/texttemplate/parse"
)

func TestAddParseTree_2(t *testing.T) {
	type args struct {
		name string
		tree *parse.Tree
	}
	tests := []struct {
		name    string
		t       *Template
		args    args
		want    *Template
		wantErr bool
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

			got, err := tt.t.AddParseTree(tt.args.name, tt.args.tree)
			if (err != nil) != tt.wantErr {
				t.Errorf("Template.AddParseTree() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Template.AddParseTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
