package orm

import (
	"fmt"
	"testing"
)

func TestLimit_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	qb := &MySQLQueryBuilder{}
	qb.Limit(10)
	if qb.String() != "LIMIT 10" {
		t.Errorf("Limit failed")
	}
}
