package orm

import (
	"database/sql"
	"fmt"
	"reflect"
	"testing"
)

func TestExec_7(t *testing.T) {
	type args struct {
		args []interface{}
	}
	tests := []struct {
		name    string
		d       *stmtQueryLog
		args    args
		want    sql.Result
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

			got, err := tt.d.Exec(tt.args.args...)
			if (err != nil) != tt.wantErr {
				t.Errorf("stmtQueryLog.Exec() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("stmtQueryLog.Exec() = %v, want %v", got, tt.want)
			}
		})
	}
}
