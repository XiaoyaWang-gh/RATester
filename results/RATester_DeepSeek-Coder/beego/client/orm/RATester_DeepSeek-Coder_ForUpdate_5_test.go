package orm

import (
	"fmt"
	"testing"
)

func TestForUpdate_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	qb := &MySQLQueryBuilder{}
	qb.ForUpdate()

	expected := " FOR UPDATE"
	actual := qb.tokens[len(qb.tokens)-1]

	if actual != expected {
		t.Errorf("Expected '%s', got '%s'", expected, actual)
	}
}
