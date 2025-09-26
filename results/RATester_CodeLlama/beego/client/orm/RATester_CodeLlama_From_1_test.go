package orm

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFrom_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	qb := &MySQLQueryBuilder{}
	qb.From("table1", "table2")
	assert.Equal(t, "FROM table1,table2", qb.String())
}
