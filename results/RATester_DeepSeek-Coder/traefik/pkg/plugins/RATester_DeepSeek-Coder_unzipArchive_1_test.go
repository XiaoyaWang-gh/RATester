package plugins

import (
	"fmt"
	"testing"
)

func TestUnzipArchive_1(t *testing.T) {
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
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			c := &Client{}
			if err := c.unzipArchive(tt.args.pName, tt.args.pVersion); (err != nil) != tt.wantErr {
				t.Errorf("Client.unzipArchive() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
