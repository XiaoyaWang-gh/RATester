package orm

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestInnerJoin_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	qb := &PostgresQueryBuilder{}
	qb.InnerJoin("table")
	assert.Equal(t, []string{"INNER JOIN", "\"table\""}, qb.tokens)
}
