package modules

import (
	"fmt"
	"testing"
)

func TestGet_11(t *testing.T) {
	type args struct {
		args []string
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

			c := &Client{
				fs:     tt.c.fs,
				logger: tt.c.logger,
			}
			if err := c.get(tt.args.args...); (err != nil) != tt.wantErr {
				t.Errorf("Client.get() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
