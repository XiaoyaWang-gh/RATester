package orm

import (
	"database/sql"
	"fmt"
	"reflect"
	"sync"
	"testing"
)

func TestGetStmt_1(t *testing.T) {
	type fields struct {
		wg   sync.WaitGroup
		stmt *sql.Stmt
	}
	tests := []struct {
		name   string
		fields fields
		want   *sql.Stmt
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

			s := &stmtDecorator{
				wg:   tt.fields.wg,
				stmt: tt.fields.stmt,
			}
			if got := s.getStmt(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("stmtDecorator.getStmt() = %v, want %v", got, tt.want)
			}
		})
	}
}
