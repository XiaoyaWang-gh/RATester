package identity

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestNewGlobIdentity_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	pattern := "*.go"
	identity := NewGlobIdentity(pattern)
	assert.NotNil(t, identity)
}
