package web

import (
	"fmt"
	"reflect"
	"testing"
)

func TestNewAppConfig_1(t *testing.T) {
	type args struct {
		appConfigProvider string
		appConfigPath     string
	}
	tests := []struct {
		name    string
		args    args
		want    *beegoAppConfig
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

			got, err := newAppConfig(tt.args.appConfigProvider, tt.args.appConfigPath)
			if (err != nil) != tt.wantErr {
				t.Errorf("newAppConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newAppConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
