package web

import (
	"fmt"
	"testing"
)

func TestParseConfig_1(t *testing.T) {
	type args struct {
		appConfigPath string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test Case 1",
			args: args{
				appConfigPath: "test_path",
			},
			wantErr: false,
		},
		{
			name: "Test Case 2",
			args: args{
				appConfigPath: "test_path_2",
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

			err := parseConfig(tt.args.appConfigPath)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
