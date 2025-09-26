package orm

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestLeftJoin_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	qb := &PostgresQueryBuilder{}
	qb.LeftJoin("table")
	assert.Equal(t, []string{"LEFT JOIN", "table"}, qb.tokens)
}
