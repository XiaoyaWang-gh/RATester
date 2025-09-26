package orm

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestOrderBy_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	qb := &PostgresQueryBuilder{}
	qb.OrderBy("a", "b")
	assert.Equal(t, "ORDER BY a, b", qb.String())
}
