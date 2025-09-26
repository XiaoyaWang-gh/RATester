package client

import (
	"fmt"
	"testing"

	v1 "github.com/fatedier/frp/pkg/config/v1"
)

func TestUpdateAllConfigurer_1(t *testing.T) {
	type args struct {
		proxyCfgs   []v1.ProxyConfigurer
		visitorCfgs []v1.VisitorConfigurer
	}
	tests := []struct {
		name    string
		svr     *Service
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

			if err := tt.svr.UpdateAllConfigurer(tt.args.proxyCfgs, tt.args.visitorCfgs); (err != nil) != tt.wantErr {
				t.Errorf("Service.UpdateAllConfigurer() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
