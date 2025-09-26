package gin

import (
	"fmt"
	"reflect"
	"testing"
)

func TestreturnObj_1(t *testing.T) {
	type fields struct {
		Handlers HandlersChain
		basePath string
		engine   *Engine
		root     bool
	}
	tests := []struct {
		name   string
		fields fields
		want   IRoutes
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

			group := &RouterGroup{
				Handlers: tt.fields.Handlers,
				basePath: tt.fields.basePath,
				engine:   tt.fields.engine,
				root:     tt.fields.root,
			}
			if got := group.returnObj(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RouterGroup.returnObj() = %v, want %v", got, tt.want)
			}
		})
	}
}
