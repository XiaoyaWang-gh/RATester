package partials

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestKey_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	k := partialCacheKey{
		Name: "test",
	}
	assert.Equal(t, "test", k.Key())
}
