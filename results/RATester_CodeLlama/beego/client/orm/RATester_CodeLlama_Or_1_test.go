package orm

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestOr_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	qb := &MySQLQueryBuilder{}
	qb.Or("cond")
	assert.Equal(t, "OR", qb.tokens[0])
	assert.Equal(t, "cond", qb.tokens[1])
}
