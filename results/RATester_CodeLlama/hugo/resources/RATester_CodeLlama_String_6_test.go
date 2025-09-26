package resources

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestString_6(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	l := &genericResource{
		name: "test",
	}
	assert.Equal(t, "Resource(genericResource: test)", l.String())
}
