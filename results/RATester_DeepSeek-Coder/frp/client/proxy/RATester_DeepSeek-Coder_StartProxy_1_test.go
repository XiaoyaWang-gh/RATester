package proxy

import (
	"fmt"
	"testing"
)

func TestStartProxy_1(t *testing.T) {
	type args struct {
		name          string
		remoteAddr    string
		serverRespErr string
	}
	tests := []struct {
		name    string
		pm      *Manager
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

			err := tt.pm.StartProxy(tt.args.name, tt.args.remoteAddr, tt.args.serverRespErr)
			if (err != nil) != tt.wantErr {
				t.Errorf("Manager.StartProxy() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
