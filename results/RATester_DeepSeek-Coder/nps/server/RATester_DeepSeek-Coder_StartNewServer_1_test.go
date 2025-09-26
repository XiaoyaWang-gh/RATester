package server

import (
	"fmt"
	"testing"

	"ehang.io/nps/lib/file"
)

func TestStartNewServer_1(t *testing.T) {
	type args struct {
		bridgePort       int
		cnf              *file.Tunnel
		bridgeType       string
		bridgeDisconnect int
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

			StartNewServer(tt.args.bridgePort, tt.args.cnf, tt.args.bridgeType, tt.args.bridgeDisconnect)
		})
	}
}
