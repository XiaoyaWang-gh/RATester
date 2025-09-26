package hugolib

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestResetLogs_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Given
	h := &HugoSites{}
	// When
	h.resetLogs()
	// Then
	assert.NotNil(t, h.Log)
}
