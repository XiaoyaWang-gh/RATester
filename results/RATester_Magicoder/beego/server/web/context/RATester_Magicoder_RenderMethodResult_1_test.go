package context

import (
	"fmt"
	"testing"
)

func TestRenderMethodResult_1(t *testing.T) {
	type args struct {
		result interface{}
	}
	tests := []struct {
		name string
		args args
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

			ctx := &Context{}
			ctx.RenderMethodResult(tt.args.result)
		})
	}
}
