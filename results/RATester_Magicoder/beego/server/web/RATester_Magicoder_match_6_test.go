package web

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/beego/beego/v2/server/web/context"
)

func Testmatch_6(t *testing.T) {
	tree := &Tree{}
	ctx := &context.Context{}

	tests := []struct {
		name           string
		treePattern    string
		pattern        string
		wildcardValues []string
		ctx            *context.Context
		wantRunObject  interface{}
	}{
		{
			name:           "Test case 1",
			treePattern:    "/test",
			pattern:        "/test",
			wildcardValues: []string{},
			ctx:            ctx,
			wantRunObject:  nil,
		},
		{
			name:           "Test case 2",
			treePattern:    "/test/:id",
			pattern:        "/test/123",
			wildcardValues: []string{"123"},
			ctx:            ctx,
			wantRunObject:  nil,
		},
		{
			name:           "Test case 3",
			treePattern:    "/test/:id",
			pattern:        "/test/456",
			wildcardValues: []string{"456"},
			ctx:            ctx,
			wantRunObject:  nil,
		},
		// Add more test cases as needed
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if gotRunObject := tree.match(tt.treePattern, tt.pattern, tt.wildcardValues, tt.ctx); !reflect.DeepEqual(gotRunObject, tt.wantRunObject) {
				t.Errorf("match() = %v, want %v", gotRunObject, tt.wantRunObject)
			}
		})
	}
}
