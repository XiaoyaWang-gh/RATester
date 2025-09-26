package orm

import (
	"fmt"
	"testing"
)

func TestAnd_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	qb := &MySQLQueryBuilder{}
	expected := "SELECT * FROM users WHERE age > 18 AND name = 'John'"
	qb.Select("*").From("users").Where("age > 18").And("name = 'John'")
	result := qb.String()
	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
}
