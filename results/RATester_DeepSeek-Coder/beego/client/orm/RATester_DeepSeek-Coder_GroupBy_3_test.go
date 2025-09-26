package orm

import (
	"fmt"
	"strings"
	"testing"
)

func TestGroupBy_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	qb := &MySQLQueryBuilder{}
	fields := []string{"field1", "field2"}
	expected := "GROUP BY field1, field2"

	qb.GroupBy(fields...)

	result := qb.String()
	if !strings.Contains(result, expected) {
		t.Errorf("Expected '%s', got '%s'", expected, result)
	}
}
