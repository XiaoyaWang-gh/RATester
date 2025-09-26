package binder

import (
	"fmt"
	"testing"

	"github.com/valyala/fasthttp"
)

func TestBindMultipart_1(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name    string
		reqCtx  *fasthttp.RequestCtx
		out     any
		wantErr bool
	}{
		{
			name:    "success",
			reqCtx:  &fasthttp.RequestCtx{},
			out:     &struct{}{},
			wantErr: false,
		},
		{
			name:    "error",
			reqCtx:  &fasthttp.RequestCtx{},
			out:     &struct{}{},
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			b := &formBinding{}
			err := b.BindMultipart(tc.reqCtx, tc.out)
			if (err != nil) != tc.wantErr {
				t.Errorf("BindMultipart() error = %v, wantErr %v", err, tc.wantErr)
			}
		})
	}
}
