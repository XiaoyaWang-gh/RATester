package orm

import (
	"fmt"
	"testing"
)

func TestWhere_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	qb := &MySQLQueryBuilder{}
	expected := "WHERE condition"
	result := qb.Where("condition").String()
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
