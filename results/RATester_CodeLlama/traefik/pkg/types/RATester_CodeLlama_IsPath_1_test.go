package types

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestIsPath_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	f := FileOrContent("path")
	assert.True(t, f.IsPath())
}
