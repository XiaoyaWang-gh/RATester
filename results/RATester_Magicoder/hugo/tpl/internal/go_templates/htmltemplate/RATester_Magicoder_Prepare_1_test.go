package template

import (
	"fmt"
	"reflect"
	"testing"

	template "github.com/gohugoio/hugo/tpl/internal/go_templates/texttemplate"
	"github.com/gohugoio/hugo/tpl/internal/go_templates/texttemplate/parse"
)

func TestPrepare_1(t *testing.T) {
	type fields struct {
		escapeErr error
		text      *template.Template
		Tree      *parse.Tree
		nameSpace *nameSpace
	}
	tests := []struct {
		name    string
		fields  fields
		want    *template.Template
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

			tr := &Template{
				escapeErr: tt.fields.escapeErr,
				text:      tt.fields.text,
				Tree:      tt.fields.Tree,
				nameSpace: tt.fields.nameSpace,
			}
			got, err := tr.Prepare()
			if (err != nil) != tt.wantErr {
				t.Errorf("Template.Prepare() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Template.Prepare() = %v, want %v", got, tt.want)
			}
		})
	}
}
