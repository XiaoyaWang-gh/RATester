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

	qs := querySet{}
	qs.Aggregate("dept_name,sum(salary) as total")

	if qs.aggregate != "dept_name,sum(salary) as total" {
		t.Errorf("Expected aggregate to be 'dept_name,sum(salary) as total', but got '%s'", qs.aggregate)
	}
}
