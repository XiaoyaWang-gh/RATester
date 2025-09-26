package client

import (
	"fmt"
	"net/http"
	"testing"
)

func TestSetAuthHeader_1(t *testing.T) {
	type fields struct {
		address  string
		authUser string
		authPwd  string
	}
	type args struct {
		req *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
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

			c := &Client{
				address:  tt.fields.address,
				authUser: tt.fields.authUser,
				authPwd:  tt.fields.authPwd,
			}
			c.setAuthHeader(tt.args.req)
			if got := tt.args.req.Header.Get("Authorization"); got != tt.want {
				t.Errorf("Client.setAuthHeader() = %v, want %v", got, tt.want)
			}
		})
	}
}
