package orm

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestInsertInto_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	qb := &PostgresQueryBuilder{}
	qb.InsertInto("table", "field1", "field2")
	assert.Equal(t, "INSERT INTO \"table\" (\"field1\",\"field2\")", qb.String())
}
