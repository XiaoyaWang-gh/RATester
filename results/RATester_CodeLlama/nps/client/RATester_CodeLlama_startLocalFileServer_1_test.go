package client

import (
	"fmt"
	"testing"

	"ehang.io/nps/lib/config"
	"ehang.io/nps/lib/file"
)

func TestStartLocalFileServer_1(t *testing.T) {
	type args struct {
		config *config.CommonConfig
		t      *file.Tunnel
		vkey   string
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

			startLocalFileServer(tt.args.config, tt.args.t, tt.args.vkey)
		})
	}
}
