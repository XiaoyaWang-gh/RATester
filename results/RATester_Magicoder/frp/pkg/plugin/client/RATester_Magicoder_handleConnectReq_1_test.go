package client

import (
	"fmt"
	"io"
	"net/http"
	"testing"
)

func TestHandleConnectReq_1(t *testing.T) {
	type args struct {
		req *http.Request
		rwc io.ReadWriteCloser
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
			hp.handleConnectReq(tt.args.req, tt.args.rwc)
		})
	}
}
