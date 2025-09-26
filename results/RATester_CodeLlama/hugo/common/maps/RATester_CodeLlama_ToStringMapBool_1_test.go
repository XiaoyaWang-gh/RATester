package maps

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestToStringMapBool_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	in := map[string]bool{"a": true, "b": false}
	assert.Equal(t, map[string]bool{"a": true, "b": false}, ToStringMapBool(in))
}
