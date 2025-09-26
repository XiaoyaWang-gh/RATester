package ssdb

import (
	"context"
	"fmt"
	"testing"

	"github.com/ssdb/gossdb/ssdb"
)

func TestClearAll_4(t *testing.T) {
	type fields struct {
		conn     *ssdb.Client
		conninfo []string
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
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

			rc := &Cache{
				conn:     tt.fields.conn,
				conninfo: tt.fields.conninfo,
			}
			if err := rc.ClearAll(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Cache.ClearAll() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
