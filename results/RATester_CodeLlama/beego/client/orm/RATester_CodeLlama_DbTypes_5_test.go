package orm

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDbTypes_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &dbBaseMysql{}
	assert.Equal(t, mysqlTypes, d.DbTypes())
}
