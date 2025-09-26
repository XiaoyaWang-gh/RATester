package orm

import (
	"database/sql"
	"fmt"
	"testing"
)

func Testdestroy_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	stmt := &sql.Stmt{}
	sd := &stmtDecorator{
		stmt: stmt,
	}
	sd.destroy()
	sd.wg.Add(1)
	sd.wg.Wait()
	if sd.stmt.Close() != nil {
		t.Errorf("Expected nil, got error")
	}
}
