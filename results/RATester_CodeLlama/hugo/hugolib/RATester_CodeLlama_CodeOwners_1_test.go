package hugolib

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestCodeOwners_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	p := &pageState{}
	p.codeowners = []string{"a", "b"}
	assert.Equal(t, p.CodeOwners(), []string{"a", "b"})
}
