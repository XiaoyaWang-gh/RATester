package orm

import (
	"fmt"
	"testing"
)

func TestIn_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	qb := &PostgresQueryBuilder{}
	vals := []string{"val1", "val2", "val3"}
	qb.In(vals...)
	expected := "IN (val1, val2, val3)"
	if qb.String() != expected {
		t.Errorf("Expected '%s', got '%s'", expected, qb.String())
	}
}
