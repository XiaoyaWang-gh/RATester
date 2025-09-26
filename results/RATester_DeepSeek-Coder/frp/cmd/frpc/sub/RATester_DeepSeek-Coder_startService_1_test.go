package sub

import (
	"fmt"
	"testing"

	v1 "github.com/fatedier/frp/pkg/config/v1"
)

func TestStartService_1(t *testing.T) {
	type args struct {
		cfg         *v1.ClientCommonConfig
		proxyCfgs   []v1.ProxyConfigurer
		visitorCfgs []v1.VisitorConfigurer
		cfgFile     string
	}
	tests := []struct {
		name    string
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

			if err := startService(tt.args.cfg, tt.args.proxyCfgs, tt.args.visitorCfgs, tt.args.cfgFile); (err != nil) != tt.wantErr {
				t.Errorf("startService() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
