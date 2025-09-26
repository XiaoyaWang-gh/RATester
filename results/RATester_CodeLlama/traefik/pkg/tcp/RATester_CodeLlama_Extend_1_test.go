package tcp

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestExtend_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := Chain{}
	chain := Chain{}
	c.Extend(chain)
	assert.Equal(t, c, chain)
}
