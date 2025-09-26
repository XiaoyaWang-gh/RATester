package orm

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestShowTablesQuery_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &dbBasePostgres{}
	assert.Equal(t, "SELECT table_name FROM information_schema.tables WHERE table_type = 'BASE TABLE' AND table_schema NOT IN ('pg_catalog', 'information_schema')", d.ShowTablesQuery())
}
