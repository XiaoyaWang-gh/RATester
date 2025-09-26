package client

import (
	"fmt"
	"testing"

	"github.com/valyala/fasthttp"
)

func TestparseCookiesFromResp_1(t *testing.T) {
	type args struct {
		host []byte
		path []byte
		resp *fasthttp.Response
	}
	tests := []struct {
		name string
		args args
		want []*fasthttp.Cookie
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

			cj := &CookieJar{}
			cj.parseCookiesFromResp(tt.args.host, tt.args.path, tt.args.resp)
		})
	}
}
