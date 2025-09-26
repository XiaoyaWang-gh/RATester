package orm

import (
	"fmt"
	"testing"
)

func TestInsertInto_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	qb := &PostgresQueryBuilder{}
	qb.InsertInto("users", "name", "email")
	expected := "INSERT INTO \"users\" (\"name\", \"email\")"
	result := qb.String()
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
