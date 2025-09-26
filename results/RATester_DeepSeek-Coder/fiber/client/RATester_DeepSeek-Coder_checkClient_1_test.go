package client

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/valyala/fasthttp"
)

func TestCheckClient_1(t *testing.T) {
	type fields struct {
		ctx          context.Context
		body         any
		header       *Header
		params       *QueryParam
		cookies      *Cookie
		path         *PathParam
		client       *Client
		formData     *FormData
		RawRequest   *fasthttp.Request
		url          string
		method       string
		userAgent    string
		boundary     string
		referer      string
		files        []*File
		timeout      time.Duration
		maxRedirects int
		bodyType     bodyType
	}
	tests := []struct {
		name   string
		fields fields
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

			r := &Request{
				ctx:          tt.fields.ctx,
				body:         tt.fields.body,
				header:       tt.fields.header,
				params:       tt.fields.params,
				cookies:      tt.fields.cookies,
				path:         tt.fields.path,
				client:       tt.fields.client,
				formData:     tt.fields.formData,
				RawRequest:   tt.fields.RawRequest,
				url:          tt.fields.url,
				method:       tt.fields.method,
				userAgent:    tt.fields.userAgent,
				boundary:     tt.fields.boundary,
				referer:      tt.fields.referer,
				files:        tt.fields.files,
				timeout:      tt.fields.timeout,
				maxRedirects: tt.fields.maxRedirects,
				bodyType:     tt.fields.bodyType,
			}
			r.checkClient()
		})
	}
}
