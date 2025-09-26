package orm

import (
	"fmt"
	"testing"
)

func TestOffset_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	qb := &MySQLQueryBuilder{}
	expected := "OFFSET 10"
	result := qb.Offset(10).String()
	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
}
