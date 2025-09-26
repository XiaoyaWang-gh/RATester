package orm

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestTableQuote_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	d := &dbBase{}
	assert.Equal(t, "`", d.TableQuote())
}
