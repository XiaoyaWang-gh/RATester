package client

import (
	"fmt"
	"testing"
)

func Testlogger_1(t *testing.T) {
	type args struct {
		c    *Client
		resp *Response
		req  *Request
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
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

			if err := logger(tt.args.c, tt.args.resp, tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("logger() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
