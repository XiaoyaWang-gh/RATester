package config

import (
	"fmt"
	"reflect"
	"testing"

	v1 "github.com/fatedier/frp/pkg/config/v1"
)

func TestLoadServerConfig_1(t *testing.T) {
	type args struct {
		path   string
		strict bool
	}
	tests := []struct {
		name    string
		args    args
		want    *v1.ServerConfig
		want1   bool
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

			got, got1, err := LoadServerConfig(tt.args.path, tt.args.strict)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadServerConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoadServerConfig() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("LoadServerConfig() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
