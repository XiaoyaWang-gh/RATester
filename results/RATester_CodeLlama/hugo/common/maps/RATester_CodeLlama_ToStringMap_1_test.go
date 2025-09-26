package maps

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestToStringMap_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	in := map[string]any{
		"a": 1,
		"b": 2,
	}
	m := ToStringMap(in)
	assert.Equal(t, map[string]any{
		"a": 1,
		"b": 2,
	}, m)
}
