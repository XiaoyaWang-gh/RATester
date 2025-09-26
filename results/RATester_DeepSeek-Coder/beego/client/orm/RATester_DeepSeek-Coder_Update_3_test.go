package orm

import (
	"fmt"
	"testing"
)

func TestUpdate_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	qb := &PostgresQueryBuilder{}
	qb.Update("table1", "table2")
	expected := "UPDATE table1, table2"
	actual := qb.String()
	if actual != expected {
		t.Errorf("Expected '%s', got '%s'", expected, actual)
	}
}
