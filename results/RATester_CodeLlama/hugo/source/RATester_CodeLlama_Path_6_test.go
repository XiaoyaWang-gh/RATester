package source

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestPath_6(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	fi := &File{}
	assert.Equal(t, "", fi.Path())
}
