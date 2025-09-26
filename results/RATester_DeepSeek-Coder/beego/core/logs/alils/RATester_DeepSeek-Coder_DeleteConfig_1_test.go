package alils

import (
	"fmt"
	"testing"
)

func TestDeleteConfig_1(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		p       *LogProject
		args    args
		wantErr bool
	}{
		{
			name: "TestDeleteConfig_Success",
			p: &LogProject{
				Name:            "test",
				Endpoint:        "http://localhost:8080",
				AccessKeyID:     "test",
				AccessKeySecret: "test",
			},
			args: args{
				name: "test_config",
			},
			wantErr: false,
		},
		{
			name: "TestDeleteConfig_Fail",
			p: &LogProject{
				Name:            "test",
				Endpoint:        "http://localhost:8080",
				AccessKeyID:     "test",
				AccessKeySecret: "test",
			},
			args: args{
				name: "non_exist_config",
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

			err := tt.p.DeleteConfig(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("DeleteConfig() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
