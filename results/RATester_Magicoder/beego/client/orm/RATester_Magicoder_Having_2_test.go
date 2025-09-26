package orm

import (
	"fmt"
	"testing"
)

func TestHaving_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	qb := &MySQLQueryBuilder{}
	expected := "HAVING test"
	result := qb.Having("test").String()
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
