package adaptor

import (
	"fmt"
	"testing"

	"github.com/valyala/fasthttp"
)

func TestCopyContextToFiberContext_1(t *testing.T) {
	type args struct {
		context        any
		requestContext *fasthttp.RequestCtx
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

			CopyContextToFiberContext(tt.args.context, tt.args.requestContext)
		})
	}
}
