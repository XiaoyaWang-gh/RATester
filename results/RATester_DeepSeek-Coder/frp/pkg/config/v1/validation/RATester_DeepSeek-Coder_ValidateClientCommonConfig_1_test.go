package validation

import (
	"fmt"
	"reflect"
	"testing"

	v1 "github.com/fatedier/frp/pkg/config/v1"
)

func TestValidateClientCommonConfig_1(t *testing.T) {
	type args struct {
		c *v1.ClientCommonConfig
	}
	tests := []struct {
		name    string
		args    args
		wantW   Warning
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

			gotW, err := ValidateClientCommonConfig(tt.args.c)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateClientCommonConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotW, tt.wantW) {
				t.Errorf("ValidateClientCommonConfig() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}
