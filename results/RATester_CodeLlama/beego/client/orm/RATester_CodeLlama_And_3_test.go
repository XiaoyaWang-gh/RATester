package orm

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnd_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	qb := &PostgresQueryBuilder{}
	cond := "id = 1"
	qb.And(cond)
	assert.Equal(t, []string{"AND", cond}, qb.tokens)
}
