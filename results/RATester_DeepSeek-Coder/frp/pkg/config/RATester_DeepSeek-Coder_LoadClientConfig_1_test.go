package config

import (
	"fmt"
	"reflect"
	"testing"

	v1 "github.com/fatedier/frp/pkg/config/v1"
)

func TestLoadClientConfig_1(t *testing.T) {
	type args struct {
		path   string
		strict bool
	}
	tests := []struct {
		name    string
		args    args
		want    *v1.ClientCommonConfig
		want1   []v1.ProxyConfigurer
		want2   []v1.VisitorConfigurer
		want3   bool
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

			got, got1, got2, got3, err := LoadClientConfig(tt.args.path, tt.args.strict)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadClientConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoadClientConfig() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("LoadClientConfig() got1 = %v, want %v", got1, tt.want1)
			}
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("LoadClientConfig() got2 = %v, want %v", got2, tt.want2)
			}
			if got3 != tt.want3 {
				t.Errorf("LoadClientConfig() got3 = %v, want %v", got3, tt.want3)
			}
		})
	}
}
