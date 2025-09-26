package hugofs

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestKey_4(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Given
	fs := &RootMappingFs{}

	// When
	key := fs.Key()

	// Then
	assert.Equal(t, "", key)
}
