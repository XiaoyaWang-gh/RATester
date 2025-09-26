package mysql

import (
	"context"
	"fmt"
	"testing"
)

func TestSessionInit_2(t *testing.T) {
	type args struct {
		ctx         context.Context
		maxlifetime int64
		savePath    string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "TestSessionInit",
			args: args{
				ctx:         context.Background(),
				maxlifetime: 100,
				savePath:    "/tmp",
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

			mp := &Provider{}
			if err := mp.SessionInit(tt.args.ctx, tt.args.maxlifetime, tt.args.savePath); (err != nil) != tt.wantErr {
				t.Errorf("Provider.SessionInit() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
