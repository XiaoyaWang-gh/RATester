package orm

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestInnerJoin_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	qb := &MySQLQueryBuilder{}
	qb.InnerJoin("table")
	assert.Equal(t, "INNER JOIN table", qb.String())
}
