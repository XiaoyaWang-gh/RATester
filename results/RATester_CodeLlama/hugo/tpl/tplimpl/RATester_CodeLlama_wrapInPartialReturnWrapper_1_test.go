package tplimpl

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/gohugoio/hugo/tpl/internal/go_templates/texttemplate/parse"
)

func TestWrapInPartialReturnWrapper_1(t *testing.T) {
	type args struct {
		n *parse.ListNode
	}
	tests := []struct {
		name string
		c    *templateContext
		args args
		want *parse.ListNode
	}{
		{
			name: "test1",
			c:    &templateContext{},
			args: args{n: &parse.ListNode{}},
			want: &parse.ListNode{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := tt.c.wrapInPartialReturnWrapper(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("templateContext.wrapInPartialReturnWrapper() = %v, want %v", got, tt.want)
			}
		})
	}
}
