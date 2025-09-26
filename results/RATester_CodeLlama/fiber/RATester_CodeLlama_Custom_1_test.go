package fiber

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestCustom_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Given
	b := &Bind{}
	name := "name"
	dest := "dest"

	// When
	err := b.Custom(name, dest)

	// Then
	assert.NoError(t, err)
}
