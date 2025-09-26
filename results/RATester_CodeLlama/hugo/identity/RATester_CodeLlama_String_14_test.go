package identity

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestString_14(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	im := &identityManager{name: "test"}
	assert.Equal(t, "IdentityManager(test)", im.String())
}
