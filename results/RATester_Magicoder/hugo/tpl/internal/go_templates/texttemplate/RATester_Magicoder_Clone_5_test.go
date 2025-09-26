package template

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/tpl/internal/go_templates/texttemplate/parse"
)

func TestClone_5(t *testing.T) {
	type fields struct {
		name       string
		Tree       *parse.Tree
		common     *common
		leftDelim  string
		rightDelim string
	}
	tests := []struct {
		name    string
		fields  fields
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

			tr := &Template{
				name:       tt.fields.name,
				Tree:       tt.fields.Tree,
				common:     tt.fields.common,
				leftDelim:  tt.fields.leftDelim,
				rightDelim: tt.fields.rightDelim,
			}
			got, err := tr.Clone()
			if (err != nil) != tt.wantErr {
				t.Errorf("Template.Clone() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Template.Clone() = %v, want %v", got, tt.want)
			}
		})
	}
}
