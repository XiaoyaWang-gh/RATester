package orm

import (
	"fmt"
	"testing"
)

func TestAggregate_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	qs := querySet{aggregate: ""}
	expected := "SUM(column)"
	result := qs.Aggregate(expected)
	if result.(*querySet).aggregate != expected {
		t.Errorf("Expected %s, got %s", expected, result.(*querySet).aggregate)
	}
}
