package validation

import (
	"fmt"
	"testing"

	v1 "github.com/fatedier/frp/pkg/config/v1"
)

func TestValidateVisitorBaseConfig_1(t *testing.T) {
	type args struct {
		c *v1.VisitorBaseConfig
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test case 1",
			args: args{
				c: &v1.VisitorBaseConfig{
					Name:       "test",
					ServerName: "test",
					BindPort:   10000,
				},
			},
			wantErr: false,
		},
		{
			name: "test case 2",
			args: args{
				c: &v1.VisitorBaseConfig{
					Name:       "test",
					ServerName: "test",
					BindPort:   0,
				},
			},
			wantErr: true,
		},
		{
			name: "test case 3",
			args: args{
				c: &v1.VisitorBaseConfig{
					Name:       "test",
					ServerName: "test",
					BindPort:   -1,
				},
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

			if err := validateVisitorBaseConfig(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("validateVisitorBaseConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
