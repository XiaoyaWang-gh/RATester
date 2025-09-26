package orm

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAnd_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	qb := &MySQLQueryBuilder{}
	qb.And("cond")
	assert.Equal(t, []string{"AND", "cond"}, qb.tokens)
}
