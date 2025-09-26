package client

import (
	"fmt"
	"net/http"
	"testing"
)

func TestDo_5(t *testing.T) {
	type args struct {
		req *http.Request
	}
	tests := []struct {
		name    string
		c       *Client
		args    args
		want    string
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

			got, err := tt.c.do(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.do() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Client.do() = %v, want %v", got, tt.want)
			}
		})
	}
}
