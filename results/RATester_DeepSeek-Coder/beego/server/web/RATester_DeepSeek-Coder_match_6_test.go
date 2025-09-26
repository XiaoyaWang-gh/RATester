package web

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/beego/beego/v2/server/web/context"
)

func TestMatch_6(t *testing.T) {
	type args struct {
		treePattern    string
		pattern        string
		wildcardValues []string
		ctx            *context.Context
	}
	tests := []struct {
		name          string
		t             *Tree
		args          args
		wantRunObject interface{}
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

			if gotRunObject := tt.t.match(tt.args.treePattern, tt.args.pattern, tt.args.wildcardValues, tt.args.ctx); !reflect.DeepEqual(gotRunObject, tt.wantRunObject) {
				t.Errorf("Tree.match() = %v, want %v", gotRunObject, tt.wantRunObject)
			}
		})
	}
}
