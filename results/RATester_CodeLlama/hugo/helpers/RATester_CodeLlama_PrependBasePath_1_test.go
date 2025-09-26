package helpers

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestPrependBasePath_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &PathSpec{}
	rel := "foo"
	isAbs := false
	assert.Equal(t, "foo", p.PrependBasePath(rel, isAbs))
}
