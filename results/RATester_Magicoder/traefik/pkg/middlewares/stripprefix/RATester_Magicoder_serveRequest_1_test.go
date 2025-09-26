package stripprefix

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestServeRequest_1(t *testing.T) {
	tests := []struct {
		name string
		req  *http.Request
		want *http.Request
	}{
		{
			name: "Test case 1",
			req: &http.Request{
				Header: http.Header{},
			},
			want: &http.Request{
				Header: http.Header{
					ForwardedPrefixHeader: []string{"prefix"},
				},
			},
		},
		{
			name: "Test case 2",
			req: &http.Request{
				Header: http.Header{
					ForwardedPrefixHeader: []string{"prefix"},
				},
			},
			want: &http.Request{
				Header: http.Header{
					ForwardedPrefixHeader: []string{"prefix"},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			s := &stripPrefix{}
			s.serveRequest(nil, tt.req, "prefix")

			if !reflect.DeepEqual(tt.req.Header, tt.want.Header) {
				t.Errorf("serveRequest() = %v, want %v", tt.req.Header, tt.want.Header)
			}
		})
	}
}
