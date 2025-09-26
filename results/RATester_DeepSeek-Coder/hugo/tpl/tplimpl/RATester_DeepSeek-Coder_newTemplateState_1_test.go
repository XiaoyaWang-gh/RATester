package tplimpl

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/identity"
	"github.com/gohugoio/hugo/tpl"
)

func TestNewTemplateState_1(t *testing.T) {
	type args struct {
		owner *templateState
		templ tpl.Template
		info  templateInfo
		id    identity.Identity
	}
	tests := []struct {
		name string
		args args
		want *templateState
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

			if got := newTemplateState(tt.args.owner, tt.args.templ, tt.args.info, tt.args.id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newTemplateState() = %v, want %v", got, tt.want)
			}
		})
	}
}
