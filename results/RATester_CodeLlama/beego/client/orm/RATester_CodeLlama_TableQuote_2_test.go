package orm

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestTableQuote_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &dbBasePostgres{}
	assert.Equal(t, `"`, d.TableQuote())
}
