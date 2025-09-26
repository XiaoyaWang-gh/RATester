package orm

import (
	"fmt"
	"testing"
)

func TestForUpdate_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	qb := &PostgresQueryBuilder{}
	qb.ForUpdate()
	expected := "FOR UPDATE"
	actual := qb.String()
	if actual != expected {
		t.Errorf("Expected '%s', got '%s'", expected, actual)
	}
}
