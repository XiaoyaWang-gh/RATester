package orm

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestGroupBy_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	qb := &PostgresQueryBuilder{}
	qb.GroupBy("a", "b")
	assert.Equal(t, "GROUP BY a, b", qb.String())
}
