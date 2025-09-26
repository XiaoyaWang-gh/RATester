package orm

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestRightJoin_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	qb := &PostgresQueryBuilder{}
	qb.RightJoin("table")
	assert.Equal(t, []string{"RIGHT JOIN", "\"table\""}, qb.tokens)
}
