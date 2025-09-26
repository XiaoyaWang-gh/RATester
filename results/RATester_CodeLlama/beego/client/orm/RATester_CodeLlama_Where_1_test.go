package orm

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestWhere_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	qb := &PostgresQueryBuilder{}
	qb.Where("id = 1")
	assert.Equal(t, "WHERE id = 1", qb.String())
}
