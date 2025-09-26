package client

import (
	"fmt"
	"testing"
)

func TestparserRequestURL_1(t *testing.T) {
	type args struct {
		c   *Client
		req *Request
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

			if err := parserRequestURL(tt.args.c, tt.args.req); (err != nil) != tt.wantErr {
				t.Errorf("parserRequestURL() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
