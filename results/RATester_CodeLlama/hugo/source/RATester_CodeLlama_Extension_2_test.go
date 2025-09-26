package source

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestExtension_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fi := &File{}
	assert.Equal(t, fi.Ext(), fi.Extension())
}
