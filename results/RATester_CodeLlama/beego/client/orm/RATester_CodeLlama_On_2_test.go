package orm

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestOn_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	qb := &PostgresQueryBuilder{}
	cond := "a.id = b.id"
	qb.On(cond)
	assert.Equal(t, []string{"ON", `"a"."id" = "b"."id"`}, qb.tokens)
}
