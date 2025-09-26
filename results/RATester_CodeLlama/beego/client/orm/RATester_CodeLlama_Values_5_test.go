package orm

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestValues_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	qb := &PostgresQueryBuilder{}
	qb.Values("a", "b", "c")
	assert.Equal(t, "VALUES (a, b, c)", qb.String())
}
