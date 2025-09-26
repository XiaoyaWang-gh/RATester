package migration

import (
	"fmt"
	"testing"
)

func TestSQL_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	m := &Migration{}
	sql := "CREATE TABLE test (id INT)"
	m.SQL(sql)
	if len(m.sqls) != 1 {
		t.Errorf("Expected 1 SQL statement, got %d", len(m.sqls))
	}
	if m.sqls[0] != sql {
		t.Errorf("Expected SQL statement to be '%s', got '%s'", sql, m.sqls[0])
	}
}
