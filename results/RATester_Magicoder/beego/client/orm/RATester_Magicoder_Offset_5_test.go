package orm

import (
	"fmt"
	"testing"
)

func TestOffset_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	qb := &PostgresQueryBuilder{}
	offset := 10
	expected := "OFFSET 10"
	qb.Offset(offset)
	result := qb.String()
	if result != expected {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}
