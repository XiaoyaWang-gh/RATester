package orm

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"testing"
)

func TestQueryRowContext_4(t *testing.T) {
	type args struct {
		ctx  context.Context
		args []interface{}
	}
	tests := []struct {
		name string
		args args
		want *sql.Row
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

			d := &stmtQueryLog{}
			if got := d.QueryRowContext(tt.args.ctx, tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("stmtQueryLog.QueryRowContext() = %v, want %v", got, tt.want)
			}
		})
	}
}
