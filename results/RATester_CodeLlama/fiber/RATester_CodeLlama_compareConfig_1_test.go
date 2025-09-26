package fiber

import (
	"fmt"
	"testing"

	"github.com/zeebo/assert"
)

func TestCompareConfig_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Given
	sf := &sendFileStore{}
	cfg := SendFile{}
	// When
	result := sf.compareConfig(cfg)
	// Then
	assert.True(t, result)
}
