package validation

import (
	"fmt"
	"reflect"
	"testing"

	v1 "github.com/fatedier/frp/pkg/config/v1"
)

func TestValidateAllClientConfig_1(t *testing.T) {
	type args struct {
		c           *v1.ClientCommonConfig
		proxyCfgs   []v1.ProxyConfigurer
		visitorCfgs []v1.VisitorConfigurer
	}
	tests := []struct {
		name    string
		args    args
		want    Warning
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

			got, err := ValidateAllClientConfig(tt.args.c, tt.args.proxyCfgs, tt.args.visitorCfgs)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateAllClientConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ValidateAllClientConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
