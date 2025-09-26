package orm

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestRightJoin_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	qb := &MySQLQueryBuilder{}
	qb.RightJoin("table")
	assert.Equal(t, "RIGHT JOIN table", qb.String())
}
