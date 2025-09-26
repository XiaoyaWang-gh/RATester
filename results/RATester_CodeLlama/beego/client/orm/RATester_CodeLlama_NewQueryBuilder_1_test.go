package orm

import (
	"fmt"
	"testing"
)

func TestNewQueryBuilder_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	var err error
	var qb QueryBuilder
	driver := "mysql"
	qb, err = NewQueryBuilder(driver)
	if err != nil {
		t.Errorf("NewQueryBuilder failed, err: %v", err)
		return
	}
	if qb == nil {
		t.Errorf("NewQueryBuilder failed, qb is nil")
		return
	}
	if qb.String() == "" {
		t.Errorf("NewQueryBuilder failed, qb.String() is empty")
		return
	}
}
