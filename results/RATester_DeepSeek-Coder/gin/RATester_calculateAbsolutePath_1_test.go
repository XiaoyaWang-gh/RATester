package gin

import (
	"fmt"
	"testing"
)

func TestCalculateAbsolutePath_1(t *testing.T) {
	type fields struct {
		Handlers HandlersChain
		basePath string
		engine   *Engine
		root     bool
	}
	type args struct {
		relativePath string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
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
			if got := group.calculateAbsolutePath(tt.args.relativePath); got != tt.want {
				t.Errorf("RouterGroup.calculateAbsolutePath() = %v, want %v", got, tt.want)
			}
		})
	}
}
