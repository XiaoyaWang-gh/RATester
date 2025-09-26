package config

import (
	"fmt"
	"reflect"
	"testing"

	"ehang.io/nps/lib/file"
)

func TestNewConfig_1(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    *Config
		wantErr bool
	}{
		{
			name: "test case 1",
			args: args{
				path: "testdata/config.conf",
			},
			want: &Config{
				content:      "testdata/config.conf",
				title:        []string{"[common]", "[host1]", "[host2]", "[tunnel1]", "[tunnel2]", "[health1]", "[health2]", "[local1]", "[local2]"},
				CommonConfig: &CommonConfig{},
				Hosts:        []*file.Host{},
				Tasks:        []*file.Tunnel{},
				Healths:      []*file.Health{},
				LocalServer:  []*LocalServer{},
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

			got, err := NewConfig(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewConfig() got = %v, want %v", got, tt.want)
			}
		})
	}
}
