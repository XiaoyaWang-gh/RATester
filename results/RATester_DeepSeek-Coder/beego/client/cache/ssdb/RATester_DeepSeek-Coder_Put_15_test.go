package ssdb

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/ssdb/gossdb/ssdb"
)

func TestPut_15(t *testing.T) {
	t.Parallel()
	ctx := context.Background()
	rc := &Cache{
		conn: &ssdb.Client{},
	}
	type args struct {
		ctx     context.Context
		key     string
		val     interface{}
		timeout time.Duration
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "normal case",
			args: args{
				ctx:     ctx,
				key:     "test_key",
				val:     "test_val",
				timeout: 0,
			},
			wantErr: false,
		},
		{
			name: "invalid value",
			args: args{
				ctx:     ctx,
				key:     "test_key",
				val:     123,
				timeout: 0,
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

			if err := rc.Put(tt.args.ctx, tt.args.key, tt.args.val, tt.args.timeout); (err != nil) != tt.wantErr {
				t.Errorf("Put() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
