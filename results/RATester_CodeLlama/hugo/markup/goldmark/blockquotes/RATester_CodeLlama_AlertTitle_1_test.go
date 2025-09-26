package blockquotes

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestAlertTitle_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &blockquoteContext{
		alert: blockQuoteAlert{
			title: "title",
		},
	}
	assert.Equal(t, "title", c.AlertTitle())
}
