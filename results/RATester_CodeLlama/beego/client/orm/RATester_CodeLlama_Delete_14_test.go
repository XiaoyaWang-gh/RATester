package orm

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDelete_14(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	qb := &MySQLQueryBuilder{}
	qb.Delete("table1", "table2")
	assert.Equal(t, "DELETE table1,table2", qb.String())
}
