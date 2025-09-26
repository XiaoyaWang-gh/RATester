package binder

import (
	"fmt"
	"testing"

	"github.com/valyala/fasthttp"
)

func TestBindMultipart_1(t *testing.T) {
	type args struct {
		reqCtx *fasthttp.RequestCtx
		out    any
	}
	tests := []struct {
		name    string
		b       *formBinding
		args    args
		wantErr bool
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

			if err := tt.b.BindMultipart(tt.args.reqCtx, tt.args.out); (err != nil) != tt.wantErr {
				t.Errorf("formBinding.BindMultipart() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
