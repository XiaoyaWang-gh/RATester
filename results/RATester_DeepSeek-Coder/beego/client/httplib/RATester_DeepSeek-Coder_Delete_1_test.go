package httplib

import (
	"fmt"
	"testing"
)

func TestDelete_1(t *testing.T) {
	type args struct {
		value interface{}
		path  string
		opts  []BeegoHTTPRequestOption
	}
	tests := []struct {
		name    string
		c       *Client
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

			if err := tt.c.Delete(tt.args.value, tt.args.path, tt.args.opts...); (err != nil) != tt.wantErr {
				t.Errorf("Client.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
