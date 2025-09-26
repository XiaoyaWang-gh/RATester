package httplib

import (
	"fmt"
	"io"
	"net/http"
	"reflect"
	"testing"
)

func TestBuildTrans_1(t *testing.T) {
	type fields struct {
		url      string
		req      *http.Request
		params   map[string][]string
		files    map[string]string
		setting  BeegoHTTPSettings
		resp     *http.Response
		body     []byte
		copyBody func() io.ReadCloser
	}
	tests := []struct {
		name   string
		fields fields
		want   http.RoundTripper
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

			b := &BeegoHTTPRequest{
				url:      tt.fields.url,
				req:      tt.fields.req,
				params:   tt.fields.params,
				files:    tt.fields.files,
				setting:  tt.fields.setting,
				resp:     tt.fields.resp,
				body:     tt.fields.body,
				copyBody: tt.fields.copyBody,
			}
			if got := b.buildTrans(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BeegoHTTPRequest.buildTrans() = %v, want %v", got, tt.want)
			}
		})
	}
}
