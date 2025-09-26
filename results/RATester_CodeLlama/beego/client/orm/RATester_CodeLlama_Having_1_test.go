package orm

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestHaving_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	qb := &PostgresQueryBuilder{}
	qb.Having("a = 1")
	assert.Equal(t, []string{"HAVING", "a = 1"}, qb.tokens)
}
