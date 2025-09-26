package orm

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestOr_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	qb := &PostgresQueryBuilder{}
	qb.Or("cond")
	assert.Equal(t, []string{"OR", "cond"}, qb.tokens)
}
