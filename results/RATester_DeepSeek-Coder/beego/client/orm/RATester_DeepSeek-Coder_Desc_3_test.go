package orm

import (
	"fmt"
	"testing"
)

func TestDesc_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	qb := &MySQLQueryBuilder{}
	qb.Desc()

	expected := "DESC"
	actual := qb.tokens[len(qb.tokens)-1]

	if actual != expected {
		t.Errorf("Expected '%s', but got '%s'", expected, actual)
	}
}
