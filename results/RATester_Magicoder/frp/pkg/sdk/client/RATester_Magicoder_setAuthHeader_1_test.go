package client

import (
	"fmt"
	"net/http"
	"testing"
)

func TestSetAuthHeader_1(t *testing.T) {
	type args struct {
		req *http.Request
	}
	tests := []struct {
		name string
		c    *Client
		args args
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

			tt.c.setAuthHeader(tt.args.req)
		})
	}
}
