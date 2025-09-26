package orm

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestShowTablesQuery_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &dbBaseSqlite{}
	assert.Equal(t, "SELECT name FROM sqlite_master WHERE type = 'table'", d.ShowTablesQuery())
}
