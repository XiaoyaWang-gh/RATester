package orm

import (
	"fmt"
	"reflect"
	"testing"
)

func TestOr_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	qb := &MySQLQueryBuilder{}
	expected := &MySQLQueryBuilder{
		tokens: []string{"OR", "cond"},
	}

	qb.Or("cond")

	if !reflect.DeepEqual(qb, expected) {
		t.Errorf("Expected %v, got %v", expected, qb)
	}
}
