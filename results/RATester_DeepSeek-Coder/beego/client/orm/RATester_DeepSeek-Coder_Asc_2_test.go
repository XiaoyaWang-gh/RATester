package orm

import (
	"fmt"
	"testing"
)

func TestAsc_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	qb := &PostgresQueryBuilder{}
	qb.Asc()
	expected := "ASC"
	actual := qb.tokens[len(qb.tokens)-1]
	if actual != expected {
		t.Errorf("Expected '%s', got '%s'", expected, actual)
	}
}
