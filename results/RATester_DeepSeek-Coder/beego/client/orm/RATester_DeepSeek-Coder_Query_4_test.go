package orm

import (
	"database/sql"
	"fmt"
	"reflect"
	"testing"
)

func TestQuery_4(t *testing.T) {
	type args struct {
		args []interface{}
	}
	tests := []struct {
		name    string
		d       *stmtQueryLog
		args    args
		want    *sql.Rows
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

			got, err := tt.d.Query(tt.args.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("stmtQueryLog.Query() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("stmtQueryLog.Query() = %v, want %v", got, tt.want)
			}
		})
	}
}
