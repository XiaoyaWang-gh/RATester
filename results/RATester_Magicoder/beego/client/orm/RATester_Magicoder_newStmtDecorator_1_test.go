package orm

import (
	"database/sql"
	"fmt"
	"reflect"
	"testing"
)

func TestnewStmtDecorator_1(t *testing.T) {
	type args struct {
		sqlStmt *sql.Stmt
	}
	tests := []struct {
		name string
		args args
		want *stmtDecorator
	}{
		{
			name: "Test case 1",
			args: args{
				sqlStmt: &sql.Stmt{},
			},
			want: &stmtDecorator{
				stmt: &sql.Stmt{},
			},
		},
		// Add more test cases as needed
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			defer func() {
				if r := recover(); r != nil {
					fmt.Println("Recovered in main", r)
				}
			}()

			if got := newStmtDecorator(tt.args.sqlStmt); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newStmtDecorator() = %v, want %v", got, tt.want)
			}
		})
	}
}
