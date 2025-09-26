package orm

import (
	"fmt"
	"testing"
)

func TestAnd_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	qb := &PostgresQueryBuilder{}
	expected := "AND condition"
	result := qb.And("condition")
	if result.String() != expected {
		t.Errorf("Expected %s, got %s", expected, result.String())
	}
}
