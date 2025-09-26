package hugolib

import (
	"fmt"
	"testing"

	"github.com/alecthomas/assert"
)

func TestResetBuildState_1(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in main", r)
		}
	}()

	// Arrange
	p := &pageState{}

	// Act
	p.resetBuildState()

	// Assert
	assert.NotNil(t, p.Scratcher)
}
