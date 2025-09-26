package orm

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestSet_35(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	qb := &MySQLQueryBuilder{}
	qb.Set("a", "b")
	assert.Equal(t, "SET a = b", qb.String())
}
