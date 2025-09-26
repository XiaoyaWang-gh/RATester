package orm

import (
	"fmt"
	"reflect"
	"testing"
)

func TestOr_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	qb := &PostgresQueryBuilder{}
	expected := &PostgresQueryBuilder{
		tokens: []string{"OR", "condition"},
	}

	result := qb.Or("condition")

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
}
