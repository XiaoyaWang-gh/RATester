package plugins

import (
	"fmt"
	"testing"
)

func TestUnzipModule_1(t *testing.T) {
	type args struct {
		pName    string
		pVersion string
	}
	tests := []struct {
		name    string
		c       *Client
		args    args
		wantErr bool
	}{
		{
			name: "Test unzipModule with valid parameters",
			c:    &Client{},
			args: args{
				pName:    "valid_name",
				pVersion: "valid_version",
			},
			wantErr: false,
		},
		{
			name: "Test unzipModule with invalid parameters",
			c:    &Client{},
			args: args{
				pName:    "invalid_name",
				pVersion: "invalid_version",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			c := &Client{}
			if err := c.unzipModule(tt.args.pName, tt.args.pVersion); (err != nil) != tt.wantErr {
				t.Errorf("Client.unzipModule() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
