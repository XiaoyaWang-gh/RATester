package orm

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestGroupBy_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	qb := &MySQLQueryBuilder{}
	qb.GroupBy("a", "b")
	assert.Equal(t, "GROUP BY a,b", qb.String())
}
