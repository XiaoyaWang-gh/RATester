package orm

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestShowTablesQuery_6(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &dbBase{}
	assert.Panics(t, func() {
		d.ShowTablesQuery()
	})
}
