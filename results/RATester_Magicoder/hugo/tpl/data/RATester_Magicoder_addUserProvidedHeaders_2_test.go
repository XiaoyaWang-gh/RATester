package data

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/gohugoio/hugo/common/types"
)

func TestaddUserProvidedHeaders_2(t *testing.T) {
	type args struct {
		headers map[string]any
		req     *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test Case 1",
			args: args{
				headers: map[string]any{
					"Content-Type":  "application/json",
					"Authorization": "Bearer token",
				},
				req: &http.Request{
					Header: make(http.Header),
				},
			},
		},
		{
			name: "Test Case 2",
			args: args{
				headers: map[string]any{
					"Content-Type": "text/plain",
					"Accept":       "*/*",
				},
				req: &http.Request{
					Header: make(http.Header),
				},
			},
		},
		{
			name: "Test Case 3",
			args: args{
				headers: nil,
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

			addUserProvidedHeaders(tt.args.headers, tt.args.req)
			for key, val := range tt.args.headers {
				vals := types.ToStringSlicePreserveString(val)
				for _, s := range vals {
					if tt.args.req.Header.Get(key) != s {
						t.Errorf("Expected %s: %s, but got %s", key, s, tt.args.req.Header.Get(key))
					}
				}
			}
		})
	}
}
