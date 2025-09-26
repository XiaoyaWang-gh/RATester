package orm

import (
	"fmt"
	"testing"
)

func TestOn_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	qb := &MySQLQueryBuilder{}
	qb.On("cond")
	if len(qb.tokens) != 2 {
		t.Errorf("On() failed")
	}
}
