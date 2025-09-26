package source

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestString_20(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fi := &File{}
	assert.Equal(t, fi.BaseFileName(), fi.String())
}
