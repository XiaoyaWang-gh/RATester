package client

import (
	"fmt"
	"testing"

	"github.com/valyala/fasthttp"
)

func TestString_2(t *testing.T) {
	type fields struct {
		client      *Client
		request     *Request
		RawResponse *fasthttp.Response
		cookie      []*fasthttp.Cookie
	}
	tests := []struct {
		name   string
		fields fields
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

			r := &Response{
				client:      tt.fields.client,
				request:     tt.fields.request,
				RawResponse: tt.fields.RawResponse,
				cookie:      tt.fields.cookie,
			}
			if got := r.String(); got != tt.want {
				t.Errorf("Response.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
