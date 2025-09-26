package hugolib

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestAliases_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &pageMeta{}
	p.pageConfig.Aliases = []string{"a", "b"}
	assert.Equal(t, p.Aliases(), []string{"a", "b"})
}
