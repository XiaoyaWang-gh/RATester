package fiber

import (
	"fmt"
	"testing"
)

func TestRenderExtensions_1(t *testing.T) {
	type args struct {
		bind any
	}
	tests := []struct {
		name string
		c    *DefaultCtx
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

			tt.c.renderExtensions(tt.args.bind)
		})
	}
}
