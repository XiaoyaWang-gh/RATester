package hugolib

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestVersion_3(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	c := &cachedContent{}
	cp := &pageContentOutput{}
	assert.Equal(t, uint32(0), c.version(cp))
}
