package orm

import (
	"fmt"
	"testing"
)

func TestSet_34(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	qb := &PostgresQueryBuilder{}
	qb.Set("a", "b")
	expected := "SET a, b"
	if qb.String() != expected {
		t.Errorf("Expected '%s', got '%s'", expected, qb.String())
	}
}
