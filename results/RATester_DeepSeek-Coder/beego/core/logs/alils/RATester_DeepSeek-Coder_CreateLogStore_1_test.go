package alils

import (
	"fmt"
	"testing"
)

func TestCreateLogStore_1(t *testing.T) {
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
			name: "Test case 1",
			p: &LogProject{
				Name:            "test",
				Endpoint:        "http://localhost:8080",
				AccessKeyID:     "test",
				AccessKeySecret: "test",
			},
			args: args{
				name:     "test_logstore",
				ttl:      1,
				shardCnt: 3,
			},
			wantErr: false,
		},
		{
			name: "Test case 2",
			p: &LogProject{
				Name:            "test",
				Endpoint:        "http://localhost:8080",
				AccessKeyID:     "test",
				AccessKeySecret: "test",
			},
			args: args{
				name:     "",
				ttl:      1,
				shardCnt: 3,
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

			if err := tt.p.CreateLogStore(tt.args.name, tt.args.ttl, tt.args.shardCnt); (err != nil) != tt.wantErr {
				t.Errorf("LogProject.CreateLogStore() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
