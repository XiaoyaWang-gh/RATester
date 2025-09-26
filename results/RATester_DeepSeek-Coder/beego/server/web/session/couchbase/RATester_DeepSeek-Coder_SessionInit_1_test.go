package couchbase

import (
	"context"
	"fmt"
	"testing"
)

func TestSessionInit_1(t *testing.T) {
	type args struct {
		ctx         context.Context
		maxlifetime int64
		cfg         string
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Test with valid json config",
			args: args{
				ctx:         context.Background(),
				maxlifetime: 100,
				cfg:         `{"save_path":"/tmp","pool":"default","bucket":"test"}`,
			},
			wantErr: false,
		},
		{
			name: "Test with invalid json config",
			args: args{
				ctx:         context.Background(),
				maxlifetime: 100,
				cfg:         `{"save_path":"/tmp","pool":"default","bucket":"test",}`,
			},
			wantErr: true,
		},
		{
			name: "Test with old style config",
			args: args{
				ctx:         context.Background(),
				maxlifetime: 100,
				cfg:         "/tmp",
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

			cp := &Provider{}
			if err := cp.SessionInit(tt.args.ctx, tt.args.maxlifetime, tt.args.cfg); (err != nil) != tt.wantErr {
				t.Errorf("SessionInit() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
