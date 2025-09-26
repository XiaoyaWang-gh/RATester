package orm

import (
	"fmt"
	"strings"
	"testing"
)

func TestIn_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	qb := &PostgresQueryBuilder{}
	vals := []string{"val1", "val2", "val3"}
	expected := "IN (val1, val2, val3)"

	qb.In(vals...)
	result := strings.Join(qb.tokens, " ")

	if result != expected {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
}
