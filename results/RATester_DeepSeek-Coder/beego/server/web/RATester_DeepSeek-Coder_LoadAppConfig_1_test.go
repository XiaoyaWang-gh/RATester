package web

import (
	"fmt"
	"testing"
)

func TestLoadAppConfig_1(t *testing.T) {
	type args struct {
		adapterName string
		configPath  string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "TestLoadAppConfig_0",
			args: args{
				adapterName: "test",
				configPath:  "test.json",
			},
			wantErr: false,
		},
		{
			name: "TestLoadAppConfig_1",
			args: args{
				adapterName: "test",
				configPath:  "not_exist.json",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			err := LoadAppConfig(tt.args.adapterName, tt.args.configPath)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadAppConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
