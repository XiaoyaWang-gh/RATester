package config

import (
	"fmt"
	"testing"
)

func TestNewConfig_1(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "TestNewConfig_0",
			args: args{
				path: "testdata/test.conf",
			},
			wantErr: false,
		},
		{
			name: "TestNewConfig_1",
			args: args{
				path: "testdata/test1.conf",
			},
			wantErr: true,
		},
		{
			name: "TestNewConfig_2",
			args: args{
				path: "testdata/test2.conf",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			_, err := NewConfig(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
