package tables

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/gohugoio/hugo/markup/converter/hooks"
)

func TestTHead_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &tableContext{}
	assert.Equal(t, c.THead(), []hooks.TableRow{})
}
