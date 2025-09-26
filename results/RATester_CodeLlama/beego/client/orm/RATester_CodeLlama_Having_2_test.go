package orm

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestHaving_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	qb := &MySQLQueryBuilder{}
	qb.Having("a = 1")
	assert.Equal(t, "HAVING a = 1", qb.String())
}
