package orm

import (
	"fmt"
	"testing"
)

func TestOr_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	qb := &PostgresQueryBuilder{}
	expected := "OR condition"
	result := qb.Or("condition")
	if result.String() != expected {
		t.Errorf("Expected %s, got %s", expected, result.String())
	}
}
