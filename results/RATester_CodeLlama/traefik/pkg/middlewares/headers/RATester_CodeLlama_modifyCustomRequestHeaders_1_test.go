package headers

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/traefik/traefik/v3/pkg/config/dynamic"
)

func TestModifyCustomRequestHeaders_1(t *testing.T) {
	type args struct {
		req *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "case 1",
			args: args{
				req: &http.Request{
					Header: make(http.Header),
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

			s := &Header{
				headers: &dynamic.Headers{
					CustomRequestHeaders: map[string]string{
						"Host": "test",
					},
				},
			}
			s.modifyCustomRequestHeaders(tt.args.req)
			if got := tt.args.req.Header.Get("Host"); got != "test" {
				t.Errorf("modifyCustomRequestHeaders() = %v, want %v", got, "test")
			}
		})
	}
}
