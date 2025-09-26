package sub

import (
	"fmt"
	"testing"
)

func TestRunClient_1(t *testing.T) {
	type args struct {
		cfgFilePath string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test Case 1",
			args: args{
				cfgFilePath: "testdata/config.yaml",
			},
			wantErr: false,
		},
		{
			name: "Test Case 2",
			args: args{
				cfgFilePath: "testdata/config.json",
			},
			wantErr: false,
		},
		{
			name: "Test Case 3",
			args: args{
				cfgFilePath: "testdata/config.toml",
			},
			wantErr: false,
		},
		{
			name: "Test Case 4",
			args: args{
				cfgFilePath: "testdata/config.ini",
			},
			wantErr: true,
		},
		{
			name: "Test Case 5",
			args: args{
				cfgFilePath: "testdata/config.unknown",
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

			if err := runClient(tt.args.cfgFilePath); (err != nil) != tt.wantErr {
				t.Errorf("runClient() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
