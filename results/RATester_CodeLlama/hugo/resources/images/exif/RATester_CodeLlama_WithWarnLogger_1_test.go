package exif

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
	"github.com/bep/logg"
)

func TestWithWarnLogger_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Given
	var warnl logg.LevelLogger
	decoder := &Decoder{}
	// When
	WithWarnLogger(warnl)(decoder)
	// Then
	assert.Equal(t, warnl, decoder.warnl)
}
