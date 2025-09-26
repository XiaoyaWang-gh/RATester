package blockquotes

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestType_5(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &blockquoteContext{typ: "test"}
	assert.Equal(t, "test", c.Type())
}
