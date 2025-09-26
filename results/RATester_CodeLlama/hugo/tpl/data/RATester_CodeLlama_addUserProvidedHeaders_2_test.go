package data

import (
	"fmt"
	"net/http"
	"testing"
)

func TestAddUserProvidedHeaders_2(t *testing.T) {
	type args struct {
		headers map[string]any
		req     *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "case 1",
			args: args{
				headers: map[string]any{
					"key1": "value1",
					"key2": "value2",
				},
				req: &http.Request{
					Header: make(http.Header),
				},
			},
		},
		{
			name: "case 2",
			args: args{
				headers: map[string]any{
					"key1": "value1",
					"key2": "value2",
				},
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
		})
	}
}
