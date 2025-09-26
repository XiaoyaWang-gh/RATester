package memcache

import (
	"context"
	"fmt"
	"testing"
)

func TestSessionInit_13(t *testing.T) {
	type args struct {
		ctx         context.Context
		maxlifetime int64
		savePath    string
	}
	tests := []struct {
		name    string
		rp      *MemProvider
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

			rp := &MemProvider{}
			if err := rp.SessionInit(tt.args.ctx, tt.args.maxlifetime, tt.args.savePath); (err != nil) != tt.wantErr {
				t.Errorf("MemProvider.SessionInit() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
