package hugofs

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
	radix "github.com/armon/go-radix"
)

func TestFilter_2(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Given
	fs := RootMappingFs{
		rootMapToReal: radix.New(),
	}

	// When
	fs.Filter(func(m RootMapping) bool {
		return true
	})

	// Then
	assert.Equal(t, fs.rootMapToReal, radix.New())
}
