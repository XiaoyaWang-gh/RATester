package types

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestString_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	f := FileOrContent("test")
	assert.Equal(t, "test", f.String())
}
