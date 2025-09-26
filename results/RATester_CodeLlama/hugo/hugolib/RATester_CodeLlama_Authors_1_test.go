package hugolib

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestAuthors_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &pageMeta{}
	assert.Nil(t, p.Authors())
}
