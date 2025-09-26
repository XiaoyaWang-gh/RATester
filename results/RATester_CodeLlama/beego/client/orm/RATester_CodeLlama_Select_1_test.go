package orm

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestSelect_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	qb := &MySQLQueryBuilder{}
	qb.Select("a", "b")
	assert.Equal(t, "SELECT a, b", qb.String())
}
