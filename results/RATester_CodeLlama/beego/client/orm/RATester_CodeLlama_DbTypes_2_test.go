package orm

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDbTypes_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &dbBaseTidb{}
	assert.Equal(t, mysqlTypes, d.DbTypes())
}
