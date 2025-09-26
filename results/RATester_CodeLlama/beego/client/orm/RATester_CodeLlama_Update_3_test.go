package orm

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestUpdate_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	qb := &PostgresQueryBuilder{}
	qb.Update("table1", "table2")
	assert.Equal(t, "UPDATE table1, table2", qb.String())
}
