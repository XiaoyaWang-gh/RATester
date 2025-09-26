package orm

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestString_29(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	qb := &MySQLQueryBuilder{}
	qb.tokens = []string{"a", "b", "c"}
	assert.Equal(t, "a b c", qb.String())
	assert.Equal(t, 0, len(qb.tokens))
}
