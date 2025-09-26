package client

import (
	"fmt"
	"net/http"
	"testing"
)

func TestConnectHandler_2(t *testing.T) {
	type args struct {
		rw  http.ResponseWriter
		req *http.Request
	}
	tests := []struct {
		name string
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

			hp := &HTTPProxy{}
			hp.ConnectHandler(tt.args.rw, tt.args.req)
		})
	}
}
