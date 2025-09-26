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
		// TODO: Add test cases.
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
