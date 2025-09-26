package hugolib

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestForEeachPageIncludingBundledPages_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Given
	include := func(p *pageState) bool {
		return true
	}
	fn := func(p *pageState) (bool, error) {
		return false, nil
	}
	m := &pageMap{}

	// When
	err := m.forEeachPageIncludingBundledPages(include, fn)

	// Then
	assert.NoError(t, err)
}
