package maps

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestToStringMapString_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	in := map[string]string{"a": "1", "b": "2"}
	m := ToStringMapString(in)
	assert.Equal(t, m, map[string]string{"a": "1", "b": "2"})
}
