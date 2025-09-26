package config

import (
	"fmt"
	"reflect"
	"testing"

	v1 "github.com/fatedier/frp/pkg/config/v1"
)

func TestLoadAdditionalClientConfigs_1(t *testing.T) {
	type args struct {
		paths          []string
		isLegacyFormat bool
		strict         bool
	}
	tests := []struct {
		name    string
		args    args
		want    []v1.ProxyConfigurer
		want1   []v1.VisitorConfigurer
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

			got, got1, err := LoadAdditionalClientConfigs(tt.args.paths, tt.args.isLegacyFormat, tt.args.strict)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadAdditionalClientConfigs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LoadAdditionalClientConfigs() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("LoadAdditionalClientConfigs() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
