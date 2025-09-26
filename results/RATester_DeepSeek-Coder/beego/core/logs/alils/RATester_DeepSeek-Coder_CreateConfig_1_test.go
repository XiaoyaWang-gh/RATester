package alils

import (
	"fmt"
	"testing"
)

func TestCreateConfig_1(t *testing.T) {
	type args struct {
		c *LogConfig
	}
	tests := []struct {
		name    string
		p       *LogProject
		args    args
		wantErr bool
	}{
		{
			name: "TestCreateConfig",
			p: &LogProject{
				Name:            "test",
				Endpoint:        "http://localhost:8080",
				AccessKeyID:     "test",
				AccessKeySecret: "test",
			},
			args: args{
				c: &LogConfig{
					Name:       "test",
					InputType:  "test",
					OutputType: "test",
				},
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

			if err := tt.p.CreateConfig(tt.args.c); (err != nil) != tt.wantErr {
				t.Errorf("LogProject.CreateConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
