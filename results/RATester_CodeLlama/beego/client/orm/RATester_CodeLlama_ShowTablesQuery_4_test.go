package orm

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestShowTablesQuery_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &dbBaseOracle{}
	assert.Equal(t, "SELECT TABLE_NAME FROM USER_TABLES", d.ShowTablesQuery())
}
