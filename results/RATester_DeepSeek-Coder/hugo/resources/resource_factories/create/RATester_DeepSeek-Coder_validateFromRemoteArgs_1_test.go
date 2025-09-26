package create

import (
	"fmt"
	"testing"
)

func TestValidateFromRemoteArgs_1(t *testing.T) {
	type args struct {
		uri     string
		options fromRemoteOptions
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

			if err := tt.c.validateFromRemoteArgs(tt.args.uri, tt.args.options); (err != nil) != tt.wantErr {
				t.Errorf("Client.validateFromRemoteArgs() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
