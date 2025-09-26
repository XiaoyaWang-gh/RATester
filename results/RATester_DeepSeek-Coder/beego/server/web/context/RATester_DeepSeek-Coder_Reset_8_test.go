package context

import (
	"fmt"
	"net/http"
	"testing"
)

func TestReset_8(t *testing.T) {
	type args struct {
		rw http.ResponseWriter
		r  *http.Request
	}
	tests := []struct {
		name string
		ctx  *Context
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

			tt.ctx.Reset(tt.args.rw, tt.args.r)
		})
	}
}
