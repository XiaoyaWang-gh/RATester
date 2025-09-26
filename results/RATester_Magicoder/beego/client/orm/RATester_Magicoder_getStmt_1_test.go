package orm

import (
	"database/sql"
	"fmt"
	"testing"
)

func TestgetStmt_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	stmt := &sql.Stmt{}
	sd := &stmtDecorator{stmt: stmt}

	got := sd.getStmt()
	if got != stmt {
		t.Errorf("getStmt() = %v, want %v", got, stmt)
	}
}
