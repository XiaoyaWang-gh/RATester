package orm

import (
	"fmt"
	"strings"
	"testing"
)

func TestFrom_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	qb := &MySQLQueryBuilder{}
	expected := "FROM table1, table2"
	result := qb.From("table1", "table2").String()
	if !strings.Contains(result, expected) {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
}
