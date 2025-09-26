package alils

import (
	"fmt"
	"testing"
)

func TestUpdateLogStore_1(t *testing.T) {
	type args struct {
		name     string
		ttl      int
		shardCnt int
	}
	tests := []struct {
		name    string
		p       *LogProject
		args    args
		wantErr bool
	}{
		{
			name: "TestUpdateLogStore_Success",
			p: &LogProject{
				Name:            "test",
				Endpoint:        "http://localhost:8080",
				AccessKeyID:     "test",
				AccessKeySecret: "test",
			},
			args: args{
				name:     "test",
				ttl:      1,
				shardCnt: 2,
			},
			wantErr: false,
		},
		{
			name: "TestUpdateLogStore_Fail",
			p: &LogProject{
				Name:            "test",
				Endpoint:        "http://localhost:8080",
				AccessKeyID:     "test",
				AccessKeySecret: "test",
			},
			args: args{
				name:     "test",
				ttl:      1,
				shardCnt: 2,
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

			if err := tt.p.UpdateLogStore(tt.args.name, tt.args.ttl, tt.args.shardCnt); (err != nil) != tt.wantErr {
				t.Errorf("LogProject.UpdateLogStore() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
