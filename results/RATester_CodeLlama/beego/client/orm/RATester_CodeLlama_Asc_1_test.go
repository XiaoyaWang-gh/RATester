package orm

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAsc_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	qb := &MySQLQueryBuilder{}
	qb.Asc()
	assert.Equal(t, "ASC", qb.String())
}
