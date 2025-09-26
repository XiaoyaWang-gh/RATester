package orm

import (
	"fmt"
	"testing"
)

func TestOn_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	qb := &MySQLQueryBuilder{}
	expected := "ON condition"
	result := qb.On("condition").String()
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
